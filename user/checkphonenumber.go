package user

import (
    "gorm.io/gorm"
    // Assuming the error helpers are available in this package
)

// checks if phone number exists in both the unverified user table and the user table.
// Takes database instance and phone number as arguments
func CheckUserPhoneNumber(Db *gorm.DB, phoneNumber string) (*PhoneNumberExists, error) {
    var verifiedcount int64
    var unverifiedcount int64

    // 1. Check Verified User Count
    if err := Db.Model(&User{}).Where("phone_number = ?", phoneNumber).Count(&verifiedcount).Error; err != nil {
        // If an error occurs here, it's a database failure (connection, query syntax, etc.)
        return nil, ErrDBLookupFailure("Failed to count verified users.", err) // Maps to 500
    } 

    // 2. Check Unverified User Count
    if err := Db.Model(&UnverifiedUser{}).Where("phone_number = ?", phoneNumber).Count(&unverifiedcount).Error; err != nil {
        // If an error occurs here, it's a database failure (connection, query syntax, etc.)
        return nil, ErrDBLookupFailure("Failed to count unverified users.", err) // Maps to 500
    } 
    
    // Success: Return the result struct
    phoneExists:= &PhoneNumberExists{
        Verified: verifiedcount > 0,
        Unverified: unverifiedcount > 0,
    }
        
    return phoneExists, nil
}