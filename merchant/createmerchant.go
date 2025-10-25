package merchant

import (
    "gorm.io/gorm"
    "github.com/GigaDesk/eardrum-interfaces/merchant"
)

// create an unverified merchant record
func CreateMerchant(s merchant.NewMerchant, Db *gorm.DB) (merchant.Merchant, error) {
    // 1. Check if the phone number already exists
    phonenumberexists, err := CheckMerchantPhoneNumber(Db, s.GetPhoneNumber())

    if err != nil {
        // Database failure during the lookup/count operation -> 500 Internal Server Error
        return nil, ErrDBLookupFailure("Failed to check merchant phone number existence.", err) 
    }

    if phonenumberexists.Unverified {
        // Business logic conflict: User exists but needs verification -> 409 Conflict
        return nil, ErrMerchantConflict("phone number already exists but is unverified. Please proceed to verification.", nil)
    }

    if phonenumberexists.Verified {
        // Business logic conflict: User fully exists -> 409 Conflict
        return nil, ErrMerchantConflict("phone number already exists and is verified.", nil)
    }

    // 2. Create unverified merchant data
    unverifiedmerchant := &UnverifiedMerchant{
        UserName:        s.GetUserName(),
        PhoneNumber: s.GetPhoneNumber(),
        Password:    s.GetPassword(),
        AccountBalanceInCents: 0,
        MpesaNumber: s.GetMpesaNumber(),
    }

    // 3. Create an unverified merchant record in the database
    if err := Db.Create(unverifiedmerchant).Error; err != nil {
        
        // Check for specific database conflict (e.g., if MpesaNumber is unique)
        if isUniqueConstraintViolation(err) {
            // Known conflict -> 409 Conflict
            return nil, ErrMerchantConflict("A unique identifier like phone number is already registered.", err)
        }

        // All other persistence failures -> 500 Internal Server Error
        return nil, ErrDBPersistenceFailure(err)
    }

    return unverifiedmerchant, nil
}


