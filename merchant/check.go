package merchant

import "gorm.io/gorm"


// CheckMerchantPhoneNumber checks if a phone number exists in either the merchant or unverified merchant table.
func CheckMerchantPhoneNumber(Db *gorm.DB, phoneNumber string) (*UniquenessCheck, error) {
	var verifiedCount int64
	var unverifiedCount int64

	// 1. Check Verified Merchant Count
	if err := Db.Model(&Merchant{}).Where("phone_number = ?", phoneNumber).Count(&verifiedCount).Error; err != nil {
		return nil, ErrDBLookupFailure("Failed to count verified merchants by phone number.", err)
	}

	// 2. Check Unverified Merchant Count
	if err := Db.Model(&UnverifiedMerchant{}).Where("phone_number = ?", phoneNumber).Count(&unverifiedCount).Error; err != nil {
		return nil, ErrDBLookupFailure("Failed to count unverified merchants by phone number.", err)
	}

	return &UniquenessCheck{
		Exists:     (verifiedCount > 0) || (unverifiedCount > 0),
		IsVerified: verifiedCount > 0,
	}, nil
}

// CheckMerchantUserName checks if a username exists in either the merchant or unverified merchant table.
func CheckMerchantUserName(Db *gorm.DB, userName string) (*UniquenessCheck, error) {
	var verifiedCount int64
	var unverifiedCount int64

	// 1. Check Verified Merchant Count
	if err := Db.Model(&Merchant{}).Where("user_name = ?", userName).Count(&verifiedCount).Error; err != nil {
		return nil, ErrDBLookupFailure("Failed to count verified merchants by username.", err)
	}

	// 2. Check Unverified Merchant Count
	if err := Db.Model(&UnverifiedMerchant{}).Where("user_name = ?", userName).Count(&unverifiedCount).Error; err != nil {
		return nil, ErrDBLookupFailure("Failed to count unverified merchants by username.", err)
	}

	return &UniquenessCheck{
		Exists:     (verifiedCount > 0) || (unverifiedCount > 0),
		IsVerified: verifiedCount > 0,
	}, nil
}