package transaction

import (
	"errors"
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"gorm.io/gorm"
)

// Gets a transaction by its unique id
func GetTransactionWithId(Db *gorm.DB, Id int) (transaction.Transaction, error) {
    var transaction *Transaction
    //fetch the record from the database
    if err := Db.First(&transaction, Id).Error; err != nil {
        // Allow gorm.ErrRecordNotFound to be returned for 404 mapping at the service layer.
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve transaction with ID %d.", Id), err)
    }

    return transaction, nil
}

// Gets all the transactions registered in the database
func GetTransactions(Db *gorm.DB) ([]transaction.Transaction, error) {

    var transactions []*Transaction

    if err := Db.Order("created_at desc").Find(&transactions).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure("Failed to retrieve all transactions.", err)
    }

    var transactionslist []transaction.Transaction

    for _, t := range transactions {
        transactionslist = append(transactionslist, t)
    }

    return transactionslist, nil
}

// GetTransactionsForUser retrieves all transactions for a specific user.
// It leverages the foreign key relationship to filter the results.
func GetTransactionsForUser(db *gorm.DB, userID uint) ([]transaction.Transaction, error) {
    var transactions []Transaction
    // GORM automatically filters on the UserID foreign key.
    if err := db.Where("user_id = ?", userID).Order("created_at desc").Find(&transactions).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve transactions for user ID %d.", userID), err)
    }
    var transactionslist []transaction.Transaction

    for _, t := range transactions {
        transactionslist = append(transactionslist, t)
    }

    return transactionslist, nil
}

// GetTransactionsForMerchant retrieves all transactions for a specific merchant.
func GetTransactionsForMerchant(db *gorm.DB, merchantID uint) ([]transaction.Transaction, error) {
    var transactions []Transaction
    if err := db.Where("merchant_id = ?", merchantID).Order("created_at desc").Find(&transactions).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve transactions for merchant ID %d.", merchantID), err)
    }
    var transactionslist []transaction.Transaction

    for _, t := range transactions {
        transactionslist = append(transactionslist, t)
    }

    return transactionslist, nil
}
