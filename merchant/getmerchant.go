package merchant

import (
    "errors"
    "gorm.io/gorm"
    "github.com/GigaDesk/eardrum-interfaces/merchant"
    // Import your custom error helpers
)

// GetMerchantWithPhoneNumber finds a merchant by phone number.
func GetMerchantWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (merchant.Merchant, error) {
    var m *Merchant

    // Find the first merchant that matches the input phonenumber
    if err := Db.Where("phone_number = ?", PhoneNumber).First(&m).Error; err != nil {
        
        // 1. Check for the known "Not Found" condition
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // Returns 404 Not Found
            return nil, ErrMerchantNotFound("phone_number", PhoneNumber) 
        }

        // 2. All other errors (connection, query syntax, etc.) -> 500 Internal
        return nil, ErrDBLookupFailure("Failed to execute query for merchant phone number.", err)
    }

    return m, nil
}

// Gets a merchant by its unique id
func GetMerchantWithId(Db *gorm.DB, Id int) (merchant.Merchant, error) {
    var m *Merchant
    
    // Fetch the record by primary key
    if err := Db.First(&m, Id).Error; err != nil {
        
        // 1. Check for Not Found
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // Returns 404 Not Found
            return nil, ErrMerchantNotFound("ID", Id) 
        }

        // 2. All other errors -> 500 Internal
        return nil, ErrDBLookupFailure("Failed to execute query for merchant ID.", err)
    }

    return m, nil
}

// Gets all the merchants registered in the database
func GetMerchants(Db *gorm.DB) ([]merchant.Merchant, error) {

    var merchants []*Merchant

    // Find all records
    if err := Db.Find(&merchants).Error; err != nil {
        // Db.Find only returns an error on connection or query issue, not if the table is empty.
        return nil, ErrDBLookupFailure("Failed to retrieve list of all merchants.", err)
    }

    // Transform [](*Merchant) to []merchant.Merchant
    var merchantlist []merchant.Merchant
    for _, m := range merchants {
        merchantlist = append(merchantlist, m)
    }

    return merchantlist, nil
}
