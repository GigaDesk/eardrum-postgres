package transaction

import (
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"gorm.io/gorm"
)

// Gets a transaction by its unique id
func GetTransactionWithId(Db *gorm.DB, Id int) (transaction.Transaction, error) {
	var transaction *Transaction
	//fetch the record from the database
	if err := Db.First(&transaction, Id).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

// Gets all the transactions registered in the database
func GetTransactions(Db *gorm.DB) ([]transaction.Transaction, error) {

	var transactions []*Transaction

	if err := Db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	var transactionslist []transaction.Transaction

	for _, t := range transactions {
		transactionslist = append(transactionslist, t)
	}

	return transactionslist, nil
}

// GetTransactionsForUser retrieves all transactions for a specific user.
// It leverages the foreign key relationship to filter the results.
func GetTransactionsForUser(db *gorm.DB, userID uint) ([]Transaction, error) {
	var transactions []Transaction
	// GORM automatically filters on the UserID foreign key.
	if err := db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionsForShop retrieves all transactions for a specific shop.
func GetTransactionsForShop(db *gorm.DB, shopID uint) ([]Transaction, error) {
	var transactions []Transaction
	if err := db.Where("shop_id = ?", shopID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
