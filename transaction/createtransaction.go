package transaction

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/product"
	"github.com/GigaDesk/eardrum-postgres/merchant"
	"github.com/GigaDesk/eardrum-postgres/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/AlekSi/pointer"
)

// ProcessOrder handles a complete, atomic financial transaction.
// It validates credentials, processes products, updates accounts, and creates all
// necessary database records securely within a single transaction.
// checkPIN is a function passed as an argument to decouple logic.
// ProcessOrder handles a complete, atomic financial transaction.
func ProcessOrder(db *gorm.DB, merchantID uint, newTx transaction.NewProductsTransaction, checkPIN func(hashedPIN, PIN string) error) (*Transaction, error) {
    var newTransaction *Transaction
    err := db.Transaction(func(tx *gorm.DB) error {
        // =========================================================================
        // 1. USER VALIDATION & ROW LOCKING
        // =========================================================================
        var u user.User
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            Where("qr_code = ?", newTx.GetUUID()).
            First(&u).Error; err != nil {
            // User not found (e.g., UUID/QRCode invalid) -> 401 Unauthorized
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return NewUnauthorizedError("invalid credentials")
            }
            // Other DB lookup issue -> 500 Internal
            return ErrDBLookupFailure("Failed to look up user for transaction.", err)
        }

        // Ensure the provided PIN is correct.
        if err := checkPIN(pointer.GetString(u.PinCode), newTx.GetPinCode()); err != nil {
            // PIN mismatch -> 401 Unauthorized
            return NewUnauthorizedError("invalid credentials.")
        }

        // =========================================================================
        // 2. PRODUCT VALIDATION & TOTAL AMOUNT CALCULATION
        // =========================================================================
        var totalAmount uint = 0
        var purchasedRecords []Purchase
        for _, item := range newTx.GetPurchasedProducts() {
            var p product.Product
            
            if err := tx.First(&p, "id = ?", item.GetProductID()).Error; err != nil {
                // Product not found -> 404 Not Found (from transaction layer)
                if errors.Is(err, gorm.ErrRecordNotFound) {
                    return NewForbiddenFailure(fmt.Sprintf("Product with ID %d not found.", item.GetProductID()))
                }
                return ErrDBLookupFailure("Failed to look up product data.", err)
            }

            // CRITICAL: Check if the product is blocked or deleted.
            if p.Blocked || p.Deleted {
                // Product unavailable -> 403 Forbidden
                return NewForbiddenFailure(fmt.Sprintf("Product with ID %d is not available for purchase.", item.GetProductID()))
            }

            // Crucial security check: Ensure the product belongs to the correct merchant.
            if p.MerchantID != merchantID {
                // Mismatched product/merchant ID -> 403 Forbidden
                return NewForbiddenFailure(fmt.Sprintf("Product ID %d does not belong to Merchant ID %d.", item.GetProductID(), merchantID))
            }

            // Calculate cost and prepare records
            itemCost := p.PricePerUnitInCents * uint(item.GetUnitsBought())
            totalAmount += itemCost
            purchasedRecords = append(purchasedRecords, Purchase{
                ProductID:          uint(item.GetProductID()),
                UnitsBought:        uint(item.GetUnitsBought()),
                TotalAmountInCents: itemCost,
            })
        }
        
        // =========================================================================
        // 3. TRANSACTION COST & FINAL BALANCE CHECK
        // =========================================================================
        feePercentageStr := os.Getenv("TRANSACTION_FEE_PERCENT")
        feePercentage, err := strconv.ParseUint(feePercentageStr, 10, 64)
        if err != nil {
            // Configuration error -> 500 Internal Server Error
            return ErrDBPersistenceFailure(errors.New("transaction fee environment variable is not properly configured"))
        }

        transactionCost := (totalAmount * uint(feePercentage)) / 100
        finalDeduction := totalAmount + transactionCost

        // Check if the user has a sufficient balance.
		if u.AccountBalanceInCents < finalDeduction {
			// Insufficient balance -> 402 Payment Required
			return NewPaymentRequiredError("Insufficient balance to complete transaction.") 
		}

        // =========================================================================
        // 4. ACCOUNT UPDATES (LOCKING BOTH USER AND SHOP)
        // =========================================================================
        var s merchant.Merchant
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            First(&s, "id = ?", merchantID).Error; err != nil {
            // Merchant not found (shouldn't happen if merchant is logged in) -> 404 Not Found
            if errors.Is(err, gorm.ErrRecordNotFound) {
                 return NewForbiddenFailure("Target merchant not found or invalid.") // Treat as 403/Forbidden to prevent merchant enumeration
            }
            return ErrDBLookupFailure("Failed to look up merchant for transaction.", err) // 500
        }

        // Perform balance updates
        u.AccountBalanceInCents -= finalDeduction
        s.AccountBalanceInCents += totalAmount

        // Save the updated balances to the database within the transaction.
        if err := tx.Save(&u).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Save error)
        }
        if err := tx.Save(&s).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Save error)
        }

        // =========================================================================
        // 5. CREATE TRANSACTION & PURCHASE RECORDS
        // =========================================================================
        newTransaction = &Transaction{
            UserID:                 u.ID,
            MerchantID:              merchantID,
            TotalAmountInCents:     totalAmount,
            TransactionCostInCents: transactionCost,
        }
        if err := tx.Create(newTransaction).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Create error)
        }

        // Create each purchase line item
        for i := range purchasedRecords {
            purchasedRecords[i].TransactionID = newTransaction.ID
            if err := tx.Create(&purchasedRecords[i]).Error; err != nil {
                return ErrDBPersistenceFailure(err) // 500 Internal (Create error)
            }
        }

        return nil // Transaction will Commit
    })

    // If the transaction failed, the structured error is returned.
    return newTransaction, err
}

// ProcessTransaction handles a complete, atomic financial transaction for a single amount.
// It validates credentials, processes the amount, updates accounts, and creates the
// necessary database records securely within a single transaction.
// checkPIN is a function passed as an argument to decouple logic.
func ProcessTransaction(db *gorm.DB, merchantID uint, newTx transaction.NewTransaction, checkPIN func(hashedPIN, PIN string) error) (*Transaction, error) {
    // We use GORM's built-in transaction helper to ensure all operations are
    // either fully completed or fully rolled back if an error occurs.
    var newTransaction *Transaction
    err := db.Transaction(func(tx *gorm.DB) error {
        // =========================================================================
        // 1. USER VALIDATION & ROW LOCKING
        // =========================================================================
        var u user.User
        // Find the user by their UUID (QrCode) and immediately lock the row
        // for the duration of this transaction to prevent race conditions.
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            Where("qr_code = ?", newTx.GetUUID()).
            First(&u).Error; err != nil {
            // User not found (QR Code invalid) -> 401 Unauthorized
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return NewUnauthorizedError("Invalid credentials")
            }
            // Other DB error during lookup -> 500 Internal Server Error
            return ErrDBLookupFailure("Failed to look up user for transaction.", err)
        }

        // Ensure the provided PIN is correct using the injected function.
        if err := checkPIN(pointer.GetString(u.PinCode), newTx.GetPinCode()); err != nil {
            // PIN mismatch -> 401 Unauthorized
            return NewUnauthorizedError("Invalid credentials")
        }

        // =========================================================================
        // 2. TRANSACTION COST & FINAL BALANCE CHECK
        // =========================================================================
        // We use the provided total amount directly, no need to calculate from products.
        totalAmount := newTx.GetTotalAmountInCents()

        // Basic validation: Amount must be positive.
        if totalAmount <= 0 {
            // Invalid amount -> 400 Bad Request
            return ErrTransactionFailed("Transaction amount must be greater than zero.")
        }

        // Retrieve the transaction cost percentage from an environment variable.
        feePercentageStr := os.Getenv("TRANSACTION_FEE_PERCENT")
        feePercentage, err := strconv.ParseUint(feePercentageStr, 10, 64)
        if err != nil {
            // Fail the transaction if the environment variable is not set or invalid. -> 500 Internal Server Error
            return ErrDBPersistenceFailure(errors.New("transaction fee environment variable is not properly configured"))
        }

        // Calculate the transaction cost using integer math to prevent floating point inaccuracies.
        transactionCost := (totalAmount * uint(feePercentage)) / 100
        finalDeduction := totalAmount + transactionCost

        // Check if the user has a sufficient balance to cover the total amount.
        if u.AccountBalanceInCents < finalDeduction {
            // Insufficient Balance -> 402 Payment Required
            return NewPaymentRequiredError("Insufficient balance to complete transaction.")
        }

        // =========================================================================
        // 3. ACCOUNT UPDATES (LOCKING BOTH USER AND MERCHANT)
        // =========================================================================
        // Find the merchant and lock its row within this same transaction to prevent race conditions.
        var s merchant.Merchant
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            First(&s, "id = ?", merchantID).Error; err != nil {
            // Merchant not found -> 403 Forbidden (prevents enumeration) or 500
            if errors.Is(err, gorm.ErrRecordNotFound) {
                 return NewForbiddenFailure("Target merchant not found or invalid.")
            }
            return ErrDBLookupFailure("Failed to look up merchant for transaction.", err) // 500
        }

        // Deduct from the user's account and add to the shop's.
        u.AccountBalanceInCents -= finalDeduction
        s.AccountBalanceInCents += totalAmount

        // Save the updated balances to the database within the transaction.
        if err := tx.Save(&u).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Save error)
        }
        if err := tx.Save(&s).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Save error)
        }

        // =========================================================================
        // 4. CREATE TRANSACTION RECORD
        // =========================================================================
        // Create the main transaction record with all calculated final amounts.
        newTransaction = &Transaction{
            UserID:                 u.ID,
            MerchantID:                 merchantID,
            TotalAmountInCents:     totalAmount,
            TransactionCostInCents: transactionCost,
        }
        if err := tx.Create(newTransaction).Error; err != nil {
            return ErrDBPersistenceFailure(err) // 500 Internal (Create error)
        }

        // This type of transaction does not have purchase records.
        // The `db.Transaction` helper will automatically commit.
        return nil
    })

    // If the transaction failed, the structured error will be returned here.
    if err != nil {
        return nil, err
    }

    return newTransaction, nil
}

