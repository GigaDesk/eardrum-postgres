package merchant

import "gorm.io/gorm"

// checks if phone number exists in both the unverified merchant table and the merchant table.
// Takes database instance and phone number as arguments
func CheckMerchantPhoneNumber(Db *gorm.DB, phoneNumber string) (*PhoneNumberExists, error) {
	var verifiedcount int64
	var unverifiedcount int64

	if err := Db.Model(&Merchant{}).Where("phone_number = ?", phoneNumber).Count(&verifiedcount).Error; err != nil {
		return nil, err
	} 

	if err := Db.Model(&UnverifiedMerchant{}).Where("phone_number = ?", phoneNumber).Count(&unverifiedcount).Error; err != nil {
		return nil, err
	} 
	
	phoneExists:= &PhoneNumberExists{
		Verified: verifiedcount > 0,
		Unverified: unverifiedcount > 0,
	}
		
	return phoneExists, nil
}