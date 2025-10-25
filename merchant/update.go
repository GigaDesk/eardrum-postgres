package merchant

import (
	"errors"
	"gorm.io/gorm"
	"github.com/GigaDesk/eardrum-interfaces/merchant"
	// Import custom error helpers
)

// UpdatePassword updates the merchant's password.
func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (merchant.Merchant, error) {
    var m *Merchant
    
    // 1. Fetch the record
    if err := Db.First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrMerchantNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch merchant for password update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&m).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record (can often be simplified, but kept for pattern)
    if err := Db.First(&m, id).Error; err != nil {
        return nil, ErrDBLookupFailure("Failed to re-fetch merchant after password update.", err) // 500
    }

    return m, nil
}

// ---

// UpdatePinCode updates the merchant's PIN code.
func UpdatePinCode(Db *gorm.DB, encryptedpincode string, id int) (merchant.Merchant, error) {
    var m *Merchant
    
    // 1. Fetch the record
    if err := Db.First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrMerchantNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch merchant for PIN update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&m).Updates(map[string]interface{}{"pin_code": encryptedpincode}).Error; err != nil {
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record
    if err := Db.First(&m, id).Error; err != nil {
        return nil, ErrDBLookupFailure("Failed to re-fetch merchant after PIN update.", err) // 500
    }

    return m, nil
}

// ---

// UpdateMpesaNumber updates the merchant's M-Pesa number.
func UpdateMpesaNumber(Db *gorm.DB, new_mpesa_number string, id int) (merchant.Merchant, error) {
    var m *Merchant
    
    // 1. Fetch the record
    if err := Db.First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrMerchantNotFound("ID", id) // 404 Not Found
        }
        return nil, ErrDBLookupFailure("Failed to fetch merchant for M-Pesa number update.", err) // 500
    }

    // 2. Update the record
    if err := Db.Model(&m).Updates(map[string]interface{}{"mpesa_number": new_mpesa_number}).Error; err != nil {
        
        // CRITICAL CHECK: If mpesa_number is a UNIQUE field, check for conflict here.
        if isUniqueConstraintViolation(err) {
            return nil, ErrMerchantConflict("The M-Pesa number is already registered to another merchant.", err) // 409 Conflict
        }
        
        return nil, ErrDBPersistenceFailure(err) // 500 Internal Server Error
    }

    // 3. Fetch the updated record
    if err := Db.First(&m, id).Error; err != nil {
        return nil, ErrDBLookupFailure("Failed to re-fetch merchant after M-Pesa number update.", err) // 500
    }

    return m, nil
}
