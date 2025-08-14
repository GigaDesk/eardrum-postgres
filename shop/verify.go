package postgresshop

import (
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

// Transforms an unverified shop record to a verified shop record
// This function uses named return variables 'verifiedShop' and 'err'
func VerifyShop(phoneNumber string, Db *gorm.DB) (verifiedShop shop.Shop, err error) {
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

    var unverifiedshop *UnverifiedShop

    // Find the unverified shop within the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Where("phone_number = ?", phoneNumber).First(&unverifiedshop).Error; err != nil {
        // The defer will handle the rollback before this return.
        return
    }

    // transform the unverified shop model into a shop model
    verifiedShop = &Shop{
        Name:                  unverifiedshop.Name,
        PhoneNumber:           unverifiedshop.PhoneNumber,
        Password:              unverifiedshop.Password,
        AccountBalanceInCents: unverifiedshop.AccountBalanceInCents,
        PinCode:               unverifiedshop.PinCode,
        MpesaNumber:           unverifiedshop.MpesaNumber,
    }

    // Create the verified shop in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Create(verifiedShop).Error; err != nil {
        return
    }

    // Delete the unverified shop in the transaction.
    // Use '=' to assign the error to the named return variable 'err'.
    if err = tx.Delete(unverifiedshop).Error; err != nil {
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