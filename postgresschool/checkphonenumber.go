package postgresschool

import "gorm.io/gorm"

// checks if phone number exists in both the unverified school table and the school table.
// Takes database instance and phone number as arguments
func CheckSchoolPhoneNumber(Db *gorm.DB, phoneNumber string) (*PhoneNumberExists, error) {
	var verifiedcount int64
	var unverifiedcount int64

	if err := Db.Model(&School{}).Where("phone_number = ?", phoneNumber).Count(&verifiedcount).Error; err != nil {
		return nil, err
	} 

	if err := Db.Model(&UnverifiedSchool{}).Where("phone_number = ?", phoneNumber).Count(&unverifiedcount).Error; err != nil {
		return nil, err
	} 
	
	phoneExists:= &PhoneNumberExists{
		Verified: verifiedcount > 0,
		Unverified: unverifiedcount > 0,
	}
		
	return phoneExists, nil
}