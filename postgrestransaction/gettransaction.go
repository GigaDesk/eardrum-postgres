package postgrestransaction

import (
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"gorm.io/gorm"
)

//Gets a transaction by its unique id
func GetTransactionWithId(Db *gorm.DB, Id int) (transaction.Transaction, error) {
	var transaction *Transaction
	//fetch the record from the database
	if err := Db.First(&transaction, Id).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

//Gets all the transactions registered in the database
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