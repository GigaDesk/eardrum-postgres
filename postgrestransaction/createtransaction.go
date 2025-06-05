package postgrestransaction

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/postgrespurchase"
	"github.com/GigaDesk/eardrum-postgres/postgresstudent"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// creates a transaction record
func CreateTransaction(t transaction.NewTransaction, checkpin func(hashedPIN string, PIN string) error, Db *gorm.DB) (transaction.Transaction, error) {

    var transaction *Transaction

	err := Db.Transaction(func(tx *gorm.DB) error {
		var student postgresstudent.Student

		// step 1: extract the student with registration number and lock
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("registration_number = ?", t.GetRegistrationNumber()).First(&student)
		if result.Error != nil {
			return result.Error
		}

		// step 2: ensure that the provided PIN code is correct
		if err := checkpin(student.PinCode, t.GetPinCode()); err != nil {
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

		//step 4: check if the students balance is more than the total

		if student.AccountBalanceInCents < int64(total) {
			return errors.New("insufficient balance to complete transaction")
		}

		//step 5: prepare transaction data
		transaction = &Transaction{
			TotalAmountInCents:   int64(total),
			BalanceBeforeInCents: student.AccountBalanceInCents,
			BalanceAfterInCents:  student.AccountBalanceInCents - int64(total),
		}

		

		//step 6: create a transaction record in the database
		if err := Db.Create(transaction).Error; err != nil {
			return err
		}

		// step 7: deduct

		student.AccountBalanceInCents -= int64(total)

		//step 8: save updates

		result = tx.Save(&student)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// creates a deposit transaction record
func CreateDepositTransaction(registrationNumber string, amountInCents int64, Db *gorm.DB) (transaction.Transaction, error) {

    var transaction *Transaction

	err := Db.Transaction(func(tx *gorm.DB) error {
		var student postgresstudent.Student

		// step 1: extract the student with registration number and lock
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("registration_number = ?", registrationNumber).First(&student)
		if result.Error != nil {
			return result.Error
		}

		//step 2: prepare transaction data
		transaction = &Transaction{
			TotalAmountInCents:   amountInCents,
			BalanceBeforeInCents: student.AccountBalanceInCents,
			BalanceAfterInCents:  student.AccountBalanceInCents + amountInCents,
		}

		

		//step 3: create a transaction record in the database
		if err := Db.Create(transaction).Error; err != nil {
			return err
		}

		// step 4: add

		student.AccountBalanceInCents += amountInCents

		//step 5: save updates

		result = tx.Save(&student)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}