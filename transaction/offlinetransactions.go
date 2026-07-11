package transaction

import (
	pgerror "errors"
	"os"
	"sort"
	"strconv"

	"github.com/GigaDesk/eardrum-interfaces/errors"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/merchant"
	"github.com/GigaDesk/eardrum-postgres/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProcessOfflineTransactionsBatch processes a block of offline transactions atomically for multiple users.
// Either all transactions succeed, or they all fail and roll back.
func ProcessOfflineTransactionsBatch(db *gorm.DB, merchantUsername string, offlineTxs []transaction.NewOfflineTransaction) ([]*Transaction, error) {
	feePercentageStr := os.Getenv("TRANSACTION_FEE_PERCENT")
	feePercentage, err := strconv.ParseUint(feePercentageStr, 10, 64)
	if err != nil {
		err3 := errors.New(errors.EARInternalError, err)
		err3.Log()
		return nil, err3
	}

	// 1. Pre-calculate totals per user to prevent multi-update deadlocks
	type UserTotals struct {
		TotalDeduction uint
		TotalMerchantCredit uint
	}
	userDeductions := make(map[string]*UserTotals)
	var runningMerchantCredit uint

	for _, offlineTx := range offlineTxs {
		phone := offlineTx.GetPhoneNumber()
		totalAmount := offlineTx.GetTotalAmountInCents()
		
		if totalAmount <= 0 {
			err3 := errors.New(errors.EARTxAmountMustBeGreaterThanZero, pgerror.New("Transaction amount must be greater than zero."))
			err3.Log()
			return nil, err3
		}

		// --- FACE MATCH VERIFICATION ---

		u , err := user.GetUserWithPhoneNumber(db, phone)
		if err!=nil{
			return nil, err
		}

        if !u.MatchFace(offlineTx.GetFacialEmbedding(), FacialMatchThreshold) {
            err3 := errors.New(errors.EARTxInvalidAuthentication, pgerror.New("Facial mismatch for phone: "+phone))
            err3.Log()
            return nil, err3
        }

		transactionCost := (totalAmount * uint(feePercentage)) / 100
		finalDeduction := totalAmount + transactionCost

		if _, exists := userDeductions[phone]; !exists {
			userDeductions[phone] = &UserTotals{}
		}
		userDeductions[phone].TotalDeduction += finalDeduction
		userDeductions[phone].TotalMerchantCredit += totalAmount
		runningMerchantCredit += totalAmount
	}

	// 2. Sort phone numbers alphabetically to guarantee strict, deterministic row-locking order
	distinctPhones := make([]string, 0, len(userDeductions))
	for phone := range userDeductions {
		distinctPhones = append(distinctPhones, phone)
	}
	sort.Strings(distinctPhones)

	var savedTransactions []*Transaction

	// 3. Begin atomic DB transaction
	err = db.Transaction(func(tx *gorm.DB) error {
		
		// Map to keep track of loaded usernames for the final transaction record step
		phoneToUsername := make(map[string]string)

		// 4. Lock and update all unique users in strict sorted order
		for _, phone := range distinctPhones {
			var u user.User
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("phone_number = ?", phone).
				First(&u).Error; err != nil {
				if pgerror.Is(err, gorm.ErrRecordNotFound) {
					err3 := errors.New(errors.EARTxUserAccountNotFound, err)
					err3.Log()
					return err3
				}
				err3 := errors.New(errors.EARInternalError, err)
				err3.Log()
				return err3
			}

			phoneToUsername[phone] = u.UserName
			totals := userDeductions[phone]

			// Overdraft logic applied to the aggregated total for this user
			if u.AccountBalanceInCents >= totals.TotalDeduction {
				u.AccountBalanceInCents -= totals.TotalDeduction
			} else {
				remainder := totals.TotalDeduction - u.AccountBalanceInCents
				u.AccountBalanceInCents = 0
				u.OverdraftBalanceInCents += remainder
			}

			if err := tx.Save(&u).Error; err != nil {
				err3 := errors.New(errors.EARInternalError, err)
				err3.Log()
				return err3
			}
		}

		// 5. Lock and update the Merchant account
		var s merchant.Merchant
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&s, "user_name = ?", merchantUsername).Error; err != nil {
			if pgerror.Is(err, gorm.ErrRecordNotFound) {
				err3 := errors.New(errors.EARTxMerchantAccountNotFound, err)
				err3.Log()
				return err3
			}
			err3 := errors.New(errors.EARMerchantLookupFailedByUsername, err)
			err3.Log()
			return err3
		}

		s.AccountBalanceInCents += runningMerchantCredit
		if err := tx.Save(&s).Error; err != nil {
			err3 := errors.New(errors.EARInternalError, err)
			err3.Log()
			return err3
		}

		// 6. Generate historical Transaction records for the batch itemization
		for _, offlineTx := range offlineTxs {
			totalAmount := offlineTx.GetTotalAmountInCents()
			transactionCost := (totalAmount * uint(feePercentage)) / 100

			newTransaction := &Transaction{
				UserUserName:           phoneToUsername[offlineTx.GetPhoneNumber()],
				MerchantUserName:       merchantUsername,
				TotalAmountInCents:     totalAmount,
				TransactionCostInCents: transactionCost,             
			}
			newTransaction.CreatedAt = offlineTx.GetOfflineTimestamp()

			if err := tx.Create(newTransaction).Error; err != nil {
				err3 := errors.New(errors.EARInternalError, err)
				err3.Log()
				return err3
			}

			savedTransactions = append(savedTransactions, newTransaction)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return savedTransactions, nil
}