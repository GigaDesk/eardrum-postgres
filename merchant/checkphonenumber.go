package merchant

import "gorm.io/gorm"

// checks if phone number exists in both the unverified merchant table and the merchant table.
// Takes database instance and phone number as arguments
func CheckMerchantPhoneNumber(Db *gorm.DB, phoneNumber string) (*PhoneNumberExists, error) {
    var verifiedcount int64
    var unverifiedcount int64

    // 1. Check Verified Merchant Count
    if err := Db.Model(&Merchant{}).Where("phone_number = ?", phoneNumber).Count(&verifiedcount).Error; err != nil {
        // Map unexpected database errors (e.g., connection failure) to 500
        return nil, ErrDBLookupFailure("Failed to count verified merchants.", err) 
    } 

    // 2. Check Unverified Merchant Count
    if err := Db.Model(&UnverifiedMerchant{}).Where("phone_number = ?", phoneNumber).Count(&unverifiedcount).Error; err != nil {
        // Map unexpected database errors to 500
        return nil, ErrDBLookupFailure("Failed to count unverified merchants.", err)
    } 
    
    // Success: Return the result struct. No GORM error means counts are accurate.
    phoneExists:= &PhoneNumberExists{
        Verified: verifiedcount > 0,
        Unverified: unverifiedcount > 0,
    }
        
    return phoneExists, nil
}