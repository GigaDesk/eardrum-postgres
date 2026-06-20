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


// GetMerchantWithUserName finds a merchant by username.
func GetMerchantWithUserName(Db *gorm.DB, userName string) (merchant.Merchant, error) {
	var m *Merchant

	// Find the first merchant that matches the input username
	if err := Db.Where("user_name = ?", userName).First(&m).Error; err != nil {
		
		// 1. Check for the known "Not Found" condition
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Returns 404 Not Found
			return nil, ErrMerchantNotFound("user_name", userName)
		}

		// 2. All other errors (connection, query syntax, etc.) -> 500 Internal
		return nil, ErrDBLookupFailure("Failed to execute query for merchant username.", err)
	}

	return m, nil
}

// GetMerchants retrieves a paginated list of merchants registered in the database.
// - limit: The maximum number of records to return (e.g., 5)
// - offset: The number of records to skip before starting to return (e.g., 0 for the first page)
func GetMerchants(Db *gorm.DB, limit int, offset int) ([]merchant.Merchant, error) {
	var merchants []*Merchant

	// Find all records with limit and offset applied
	if err := Db.Limit(limit).Offset(offset).Find(&merchants).Error; err != nil {
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
