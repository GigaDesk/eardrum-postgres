package merchant

import (
	"time"

	"gorm.io/gorm"
)

// Merchant represents the merchant model for the system.
type Merchant struct {
    gorm.Model
    UserName              string `gorm:"not null"` // UserName of the merchant
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the merchant
    Password              string `gorm:"not null"` // The merchants's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The merchant's account balance in cents
    PinCode               *string `gorm:""` // The merchant's security PIN code. It should be stored as a secure hash (NULLABLE).
    MpesaNumber           *string `gorm:""` // The merchant's M-Pesa number used for withdrawals(NULLABLE)
}

// Returns the unique ID of the merchant
func (s Merchant) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the merchant
func (s Merchant) GetCreatedAt() time.Time {
	return s.CreatedAt.UTC()
}

// Returns the update timestamp of the merchant
func (s Merchant) GetUpdatedAt() time.Time {
	return s.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the merchant
func (s Merchant) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.UTC()
}

// Returns the name of the merchant
func (s Merchant) GetUserName() string {
	return s.UserName
}

// Returns the phone number of the merchant
func (s Merchant) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the m-pesa number of the merchant
func (s Merchant) GetMpesaNumber() *string {
	return s.MpesaNumber
}

// Returns the account balance in cents of the merchant
func (s Merchant) GetAccountBalanceInCents() int64 {
	return int64(s.AccountBalanceInCents)
}

// Returns the security password of the shop
func (s Merchant) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the shop
func (s Merchant) GetPinCode() *string {
	return s.PinCode
}

// Merchant represents the merchant model for the system.
type UnverifiedMerchant struct {
    gorm.Model
    UserName              string `gorm:"not null"` // Name of the merchant
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the merchant
    Password              string `gorm:"not null"` // The merchant's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The merchant's account balance in cents
    PinCode               *string `gorm:""` // The merchant's security PIN code. It should be stored as a secure hash.
    MpesaNumber           *string `gorm:""` // The merchant's M-Pesa number used for withdrawals
}
// Returns the unique ID of the unverified merchant
func (s UnverifiedMerchant) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the unverified merchant
func (s UnverifiedMerchant) GetCreatedAt() time.Time {
	return s.CreatedAt.UTC()
}

// Returns the update timestamp of the unverified merchant
func (s UnverifiedMerchant) GetUpdatedAt() time.Time {
	return s.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the unverified merchant
func (s UnverifiedMerchant) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.UTC()
}

// Returns the name of the unverified merchant
func (s UnverifiedMerchant) GetUserName() string {
	return s.UserName
}

// Returns the phone number of the unverified merchant
func (s UnverifiedMerchant) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the account balance in cents of the unverified merchant
func (s UnverifiedMerchant) GetAccountBalanceInCents() int64 {
	return int64(s.AccountBalanceInCents)
}

// Returns the security password of the unverified merchant
func (s UnverifiedMerchant) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the unverified merchant
func (s UnverifiedMerchant) GetPinCode() *string {
	return s.PinCode
}

// Returns the m-pesa number of the unverified merchant
func (s UnverifiedMerchant) GetMpesaNumber() *string {
	return s.MpesaNumber
}

type PhoneNumberExists struct {
	Verified   bool
	Unverified bool
}
