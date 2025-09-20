package user

import (
    "github.com/GigaDesk/eardrum-interfaces/user"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// RegenerateQrCode generates a new UUID for the user's QR code, effectively invalidating the old one.
// It uses GORM's built-in transaction function with a row-level lock for thread safety.
func RegenerateQrCode(Db *gorm.DB, id int) (user.User, error) {
    var verifiedUser *User

    err := Db.Transaction(func(tx *gorm.DB) error {
        // Find the user within the transaction and add a lock.
        // This ensures that no other process can modify this user's record
        // until the transaction is committed or rolled back.
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&verifiedUser).Error; err != nil {
            return err
        }

        // Generate a new, cryptographically strong UUID.
        newQrCode := uuid.New()

        // Update the user's record with the new UUID.
        if err := tx.Model(&verifiedUser).Update("qr_code", newQrCode).Error; err != nil {
            return err
        }

        // Update the local struct to reflect the new UUID
        verifiedUser.QrCode = newQrCode

        // Return nil to commit the transaction
        return nil
    })

    if err != nil {
        return nil, err
    }

    // The transaction was successful.
    return verifiedUser, nil
}