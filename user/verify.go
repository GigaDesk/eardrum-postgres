package user

import (
    "errors"
    "gorm.io/gorm"
    "github.com/GigaDesk/eardrum-interfaces/user"
)

// Transforms an unverified user record to a verified user record
func VerifyUser(phoneNumber string, Db *gorm.DB) (verifiedUser user.User, finalErr error) {
    
    // Start a new transaction
    tx := Db.Begin()
    if tx.Error != nil {
        // If the transaction can't even start (e.g., connection issue)
        return nil, ErrDBPersistenceFailure(tx.Error) // 500 Internal
    }

    // Defer a rollback that will only execute if an error occurred in the transaction body.
    defer func() {
        if finalErr != nil {
            tx.Rollback()
        }
    }()

    var unverifieduser *UnverifiedUser

    // 1. Find the unverified user within the transaction.
    if err := tx.Where("phone_number = ?", phoneNumber).First(&unverifieduser).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // User not found (expected business failure) -> 404 Not Found
            finalErr = ErrUserNotFound("unverified phone_number", phoneNumber)
            return 
        }
        // Other lookup failure (e.g., connection issue) -> 500 Internal
        finalErr = ErrDBLookupFailure("Failed to find unverified user.", err)
        return
    }

    // transform the unverified user model into a user model
    verifiedUser = &User{
        // ... field assignments
        UserName:              unverifieduser.UserName,
        PhoneNumber:           unverifieduser.PhoneNumber,
        Password:              unverifieduser.Password,
        AccountBalanceInCents: unverifieduser.AccountBalanceInCents,
        PinCode:               unverifieduser.PinCode,
        MpesaNumber:           unverifieduser.MpesaNumber,
        QrCode:                unverifieduser.QrCode,
    }

    // 2. Create the verified user in the transaction.
    if err := tx.Create(verifiedUser).Error; err != nil {
        // If creation fails, check for a unique constraint violation (e.g., QR Code clash)
        if isUniqueConstraintViolation(err) {
            finalErr = ErrUserConflict("Verification failed: User record already exists or unique ID is duplicated.")
            return // 409 Conflict
        }
        // Generic create failure -> 500 Internal
        finalErr = ErrDBPersistenceFailure(err)
        return
    }

    // 3. Delete the unverified user in the transaction.
    if err := tx.Delete(unverifieduser).Error; err != nil {
        // Delete failure -> 500 Internal
        finalErr = ErrDBPersistenceFailure(err)
        return
    }

    // 4. Commit the transaction.
    if err := tx.Commit().Error; err != nil {
        // Commit failure -> 500 Internal
        finalErr = ErrDBPersistenceFailure(err)
        return
    }

    // All successful, return the verified user (finalErr is nil)
    return
}
