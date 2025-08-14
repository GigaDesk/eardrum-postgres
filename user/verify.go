package user

import (
    "github.com/GigaDesk/eardrum-interfaces/user"
    "gorm.io/gorm"
)

// Transforms an unverified user record to a verified user record
// This function uses named return variables 'verifiedUser' and 'err'
func VerifyUser(phoneNumber string, Db *gorm.DB) (verifiedUser user.User, err error) {
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

    var unverifieduser *UnverifiedUser

    // Find the unverified user within the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Where("phone_number = ?", phoneNumber).First(&unverifieduser).Error; err != nil {
        // The defer will handle the rollback before this return.
        return
    }

    // transform the unverified user model into a user model
    verifiedUser = &User{
        Name:                  unverifieduser.Name,
        PhoneNumber:           unverifieduser.PhoneNumber,
        Password:              unverifieduser.Password,
        AccountBalanceInCents: unverifieduser.AccountBalanceInCents,
        PinCode:               unverifieduser.PinCode,
        MpesaNumber:           unverifieduser.MpesaNumber,
    }

    // Create the verified user in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Create(verifiedUser).Error; err != nil {
        return
    }

    // Delete the unverified user in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Delete(unverifieduser).Error; err != nil {
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