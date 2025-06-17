package postgrestransaction

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-interfaces/user"
	"github.com/GigaDesk/eardrum-postgres/postgrespurchase"
	"github.com/GigaDesk/eardrum-postgres/postgresshop"
	"github.com/GigaDesk/eardrum-postgres/postgresuser"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// deposits a particular amount to a shop's account
func AddToShopBalance(shopid int, amountInCents int64, Db *gorm.DB) (shop.Shop, error) {

	var shop postgresshop.Shop

	err := Db.Transaction(func(tx *gorm.DB) error {

		// step 1: extract the shop with the shopid and lock
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&shop, shopid)
		if result.Error != nil {
			return result.Error
		}

		// step 2: add

		shop.AccountBalanceInCents += amountInCents

		//step 3: save updates

		result = tx.Save(&shop)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return shop, nil
}

// creates a transaction record
func CreateTransaction(t transaction.NewTransaction, shopid int, checkpin func(hashedPIN string, PIN string) error, Db *gorm.DB) (transaction.Transaction, user.User, shop.Shop, []purchase.Purchase, error) {

	var transaction *Transaction

	var purchases []purchase.Purchase

	var shop shop.Shop

	var user *postgresuser.User

	err := Db.Transaction(func(tx *gorm.DB) error {

		// step 1: extract the user with phone number and lock
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("phone_number = ?", t.GetPhoneNumber()).First(&user)
		if result.Error != nil {
			return result.Error
		}

		// step 2: ensure that the provided PIN code is correct
		if err := checkpin(user.PinCode, t.GetPinCode()); err != nil {
			return err
		}

		//step 3: Calculate the total amount involved in the transaction

		var total int

		for _, p := range t.GetPurchasedProducts() {
			amount, err := postgrespurchase.CalculatePurchaseAmountInCents(p, Db)
			if err != nil {
				return err
			}
			total += amount
		}

		//step 4: check if the user's balance is more than the total

		if user.AccountBalanceInCents < int64(total) {
			return errors.New("insufficient balance to complete transaction")
		}

		//step 5: prepare transaction data
		transaction = &Transaction{
			TotalAmountInCents:     int64(total),
			TransactionCostInCents: 0,
		}

		//step 6: create a transaction record in the database
		if err := Db.Create(transaction).Error; err != nil {
			return err
		}

		// step 7: deduct

		user.AccountBalanceInCents -= int64(total)

		//step 8: save updates to user

		result = tx.Save(&user)
		if result.Error != nil {
			return result.Error
		}

		//step 9: Add to shop balance
		s, err := AddToShopBalance(shopid, int64(total), Db)

		if err != nil {
			return err
		}

		shop = s

		for _, p := range t.GetPurchasedProducts() {
			purchase, err := postgrespurchase.CreatePurchase(transaction, p, Db)
			if err != nil {
				return err
			}
			purchases = append(purchases, purchase)
		}

		return nil
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return transaction, user, shop, purchases, nil
}
