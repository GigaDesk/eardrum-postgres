package merchant

import (
	"github.com/GigaDesk/eardrum-interfaces/errors"
	"github.com/GigaDesk/eardrum-interfaces/merchant"
	"gorm.io/gorm"
)

// CreateMerchant creates an unverified merchant record after ensuring the phone number and username are unique.
func CreateMerchant(s merchant.NewMerchant, Db *gorm.DB) (merchant.Merchant, error) {
	// 1. Check if the phone number already exists
	phoneCheck, err := CheckMerchantPhoneNumber(Db, s.GetPhoneNumber())
	if err != nil {
		// Database failure during the lookup/count operation -> 500 Internal Server Error
		return nil, ErrDBLookupFailure("Failed to check merchant phone number existence.", err)
	}

	if phoneCheck.Exists {
		if !phoneCheck.IsVerified {
			// Business logic conflict: User exists but needs verification -> 409 Conflict
			return nil, ErrMerchantConflict("phone number already exists but is unverified. Please proceed to verification.", nil)
		}
		// Business logic conflict: User fully exists -> 409 Conflict
		return nil, ErrMerchantConflict("phone number already exists and is verified.", nil)
	}

	// 2. Check if the username already exists
	usernameCheck, err := CheckMerchantUserName(Db, s.GetUserName())
	if err != nil {
		// Database failure during the lookup/count operation -> 500 Internal Server Error
		return nil, ErrDBLookupFailure("Failed to check merchant username existence.", err)
	}

	if usernameCheck.Exists {
		if !usernameCheck.IsVerified {
			// Business logic conflict: Username exists in unverified state -> 409 Conflict
			return nil, ErrMerchantConflict("username already exists but is unverified.", nil)
		}
		// Business logic conflict: Username fully exists -> 409 Conflict
		return nil, ErrMerchantConflict("username already exists and is verified.", nil)
	}

	// 3. Create unverified merchant data
	unverifiedmerchant := &UnverifiedMerchant{
		UserName:              s.GetUserName(),
		PhoneNumber:           s.GetPhoneNumber(),
		Password:              s.GetPassword(),
		AccountBalanceInCents: 0,
	}

	// 4. Create an unverified merchant record in the database
	if err := Db.Create(unverifiedmerchant).Error; err != nil {

		// All other persistence failures -> 500 Internal Server Error
		err1 := errors.New(errors.EARInternalError, err)
		err1.Log()
		return nil, err1
	}

	return unverifiedmerchant, nil
}
