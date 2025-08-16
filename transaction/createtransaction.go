package transaction

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/GigaDesk/eardrum-postgres/shop"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/product"
	"github.com/GigaDesk/eardrum-postgres/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProcessOrder handles a complete, atomic financial transaction.
// It validates credentials, processes products, updates accounts, and creates all
// necessary database records securely within a single transaction.
// checkPIN is a function passed as an argument to decouple logic.
func ProcessOrder(db *gorm.DB, shopID uint, newTx transaction.NewTransaction, checkPIN func(hashedPIN, PIN string) error) (*Transaction, error) {
	// We use GORM's built-in transaction helper to ensure all operations are
	// either fully completed or fully rolled back if an error occurs.
	var newTransaction *Transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// =========================================================================
		// 1. USER VALIDATION & ROW LOCKING
		// =========================================================================
		var u user.User
		// Find the user by phone number and immediately lock the row for the duration
		// of this transaction to prevent race conditions.
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("phone_number = ?", newTx.GetPhoneNumber()).
			First(&u).Error; err != nil {
			return errors.New("invalid credentials")
		}

		// Ensure the provided PIN is correct using the injected function.
		if err := checkPIN(u.PinCode, newTx.GetPinCode()); err != nil {
			return errors.New("invalid credentials")
		}

		// =========================================================================
		// 2. PRODUCT VALIDATION & TOTAL AMOUNT CALCULATION (WITHOUT STOCK CHECK)
		// =========================================================================
		var totalAmount uint = 0
		var purchasedRecords []Purchase
		for _, item := range newTx.GetPurchasedProducts() {
			var p product.Product
			// Find the product in the database.
			if err := tx.First(&p, "id = ?", item.GetProductID()).Error; err != nil {
				return errors.New("product not found")
			}
			
			// Crucial security check: Ensure the product belongs to the correct shop.
			if p.ShopID != shopID {
				return fmt.Errorf("product ID %d does not belong to shop ID %d", item.GetProductID(), shopID)
			}
			
			// Calculate the cost for this purchase item.
			itemCost := p.PricePerUnitInCents * uint(item.GetUnitsBought())
			totalAmount += itemCost
			
			// Prepare the individual purchase record to be created later.
			newPurchase := Purchase{
				ProductID:          uint(item.GetProductID()),
				UnitsBought:        uint(item.GetUnitsBought()),
				TotalAmountInCents: itemCost,
			}
			purchasedRecords = append(purchasedRecords, newPurchase)
		}

		// =========================================================================
		// 3. TRANSACTION COST & FINAL BALANCE CHECK
		// =========================================================================
		// Retrieve the transaction cost percentage from an environment variable.
		feePercentageStr := os.Getenv("TRANSACTION_FEE_PERCENT")
		feePercentage, err := strconv.ParseUint(feePercentageStr, 10, 64)
		if err != nil {
			// Fail the transaction if the environment variable is not set or invalid.
			return errors.New("transaction fee environment variable is not properly configured")
		}

		// Calculate the transaction cost using integer math to prevent floating point inaccuracies.
		transactionCost := (totalAmount * uint(feePercentage)) / 100
		finalDeduction := totalAmount + transactionCost
		
		// Check if the user has a sufficient balance to cover the total amount.
		if u.AccountBalanceInCents < finalDeduction {
			return errors.New("insufficient balance to complete transaction")
		}

		// =========================================================================
		// 4. ACCOUNT UPDATES (LOCKING BOTH USER AND SHOP)
		// =========================================================================
		// Find the shop and lock its row within this same transaction to prevent race conditions.
		var s shop.Shop
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&s, "id = ?", shopID).Error; err != nil {
			return errors.New("shop not found or invalid")
		}
		
		// Deduct from the user's account and add to the shop's.
		u.AccountBalanceInCents -= finalDeduction
		s.AccountBalanceInCents += totalAmount
		
		// Save the updated balances to the database within the transaction.
		if err := tx.Save(&u).Error; err != nil {
			return err
		}
		if err := tx.Save(&s).Error; err != nil {
			return err
		}

		// =========================================================================
		// 5. CREATE TRANSACTION & PURCHASE RECORDS
		// =========================================================================
		// Create the main transaction record with all calculated final amounts.
		newTransaction = &Transaction{
			UserID:                 u.ID,
			ShopID:                 shopID,
			TotalAmountInCents:     totalAmount,
			TransactionCostInCents: transactionCost,
		}
		if err := tx.Create(newTransaction).Error; err != nil {
			return err
		}

		// Create each purchase line item, linking it to the new transaction ID.
		for i := range purchasedRecords {
			purchasedRecords[i].TransactionID = newTransaction.ID
			if err := tx.Create(&purchasedRecords[i]).Error; err != nil {
				return err
			}
		}

		// If we've reached this point, all operations were successful.
		// The `db.Transaction` helper will automatically commit.
		return nil
	})

	// If the transaction failed, the error will be returned here.
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

