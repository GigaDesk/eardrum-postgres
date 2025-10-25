package merchant

import (
	"github.com/GigaDesk/eardrum-interfaces/merchant"
	"gorm.io/gorm"
)

// Transforms an unverified shop record to a verified merchant record
// This function uses named return variables 'verifiedMerchant' and 'err'
func VerifyMerchant(phoneNumber string, Db *gorm.DB) (verifiedMerchant merchant.Merchant, err error) {
    // Start a new transaction
    tx := Db.Begin()
    if tx.Error != nil {
        // If the transaction can't start, we return the error immediately.
        return nil, tx.Error
    }

    // Defer a rollback that will only execute if an error occurred.
    // This closure now correctly refers to the function's named return variable 'err'.
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    var unverifiedmerchant *UnverifiedMerchant

    // Find the unverified merchant within the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Where("phone_number = ?", phoneNumber).First(&unverifiedmerchant).Error; err != nil {
        // The defer will handle the rollback before this return.
        return
    }

    // transform the unverified merchant model into a merchant model
    verifiedMerchant = &Merchant{
        UserName:              unverifiedmerchant.UserName,
        PhoneNumber:           unverifiedmerchant.PhoneNumber,
        Password:              unverifiedmerchant.Password,
        AccountBalanceInCents: unverifiedmerchant.AccountBalanceInCents,
        PinCode:               unverifiedmerchant.PinCode,
        MpesaNumber:           unverifiedmerchant.MpesaNumber,
    }

    // Create the verified merchant in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Create(verifiedMerchant).Error; err != nil {
        return
    }

    // Delete the unverified merchant in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Delete(unverifiedmerchant).Error; err != nil {
        return
    }

    // Commit the transaction if all operations were successful.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Commit().Error; err != nil {
        return
    }

    // The function returns here. The defer closure runs and finds 'err' is nil,
    // so it skips the rollback.
    return
}