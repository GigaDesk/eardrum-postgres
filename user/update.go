package user

import (
	"errors"
	"gorm.io/gorm"
	"github.com/GigaDesk/eardrum-interfaces/user"
)

func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (user.User, error) {
    var u *User
    
    // 1. Fetch the record
    if err := Db.First(&u, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch user for password update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&u).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record (can often be skipped if the data is reliable)
    if err := Db.First(&u, id).Error; err != nil {
        // Defensive check: This should not fail if the update just succeeded.
        return nil, ErrDBLookupFailure("Failed to re-fetch user after password update.", err) // 500
    }

    // Return the updated record
    return u, nil
}

func UpdatePinCode(Db *gorm.DB, encryptedpincode string, id int) (user.User, error) {
    var u *User
    
    // 1. Fetch the record
    if err := Db.First(&u, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch user for PIN update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&u).Updates(map[string]interface{}{"pin_code": encryptedpincode}).Error; err != nil {
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record
    if err := Db.First(&u, id).Error; err != nil {
        return nil, ErrDBLookupFailure("Failed to re-fetch user after PIN update.", err) // 500
    }

    return u, nil
}

func UpdateMpesaNumber(Db *gorm.DB, new_mpesa_number string, id int) (user.User, error) {
    var u *User
    
    // 1. Fetch the record
    if err := Db.First(&u, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch user for M-Pesa number update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&u).Updates(map[string]interface{}{"mpesa_number": new_mpesa_number}).Error; err != nil {
        // NOTE: If "mpesa_number" is a UNIQUE field, you should specifically check for 
        // a unique constraint violation here and map it to an ErrUserConflict (409).
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record
    if err := Db.First(&u, id).Error; err != nil {
        return nil, ErrDBLookupFailure("Failed to re-fetch user after M-Pesa number update.", err) // 500
    }

    return u, nil
}