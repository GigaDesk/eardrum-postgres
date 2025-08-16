package shop

import (
	"time"

	"gorm.io/gorm"
)

// Shop represents the shop model for the system.
type Shop struct {
    gorm.Model
    Name                  string `gorm:"not null"` // Name of the shop
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the user
    Password              string `gorm:"not null"` // The shop's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The shop's account balance in cents
    PinCode               string `gorm:"not null"` // The shop's security PIN code. It should be stored as a secure hash.
    MpesaNumber           string `gorm:"not null"` // The users M-Pesa number used for withdrawals
}

// Returns the unique ID of the shop
func (s Shop) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the shop
func (s Shop) GetCreatedAt() time.Time {
	return s.CreatedAt.UTC()
}

// Returns the update timestamp of the shop
func (s Shop) GetUpdatedAt() time.Time {
	return s.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the shop
func (s Shop) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.UTC()
}

// Returns the name of the shop
func (s Shop) GetName() string {
	return s.Name
}

// Returns the phone number of the shop
func (s Shop) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the m-pesa number of the shop
func (s Shop) GetMpesaNumber() string {
	return s.MpesaNumber
}

// Returns the account balance in cents of the shop
func (s Shop) GetAccountBalanceInCents() int64 {
	return int64(s.AccountBalanceInCents)
}

// Returns the security password of the shop
func (s Shop) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the shop
func (s Shop) GetPinCode() string {
	return s.PinCode
}

// Shop represents the shop model for the system.
type UnverifiedShop struct {
    gorm.Model
    Name                  string `gorm:"not null"` // Name of the shop
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the user
    Password              string `gorm:"not null"` // The shop's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The shop's account balance in cents
    PinCode               string `gorm:"not null"` // The shop's security PIN code. It should be stored as a secure hash.
    MpesaNumber           string `gorm:"not null"` // The users M-Pesa number used for withdrawals
}
// Returns the unique ID of the unverified shop
func (s UnverifiedShop) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the unverified shop
func (s UnverifiedShop) GetCreatedAt() time.Time {
	return s.CreatedAt.UTC()
}

// Returns the update timestamp of the unverified shop
func (s UnverifiedShop) GetUpdatedAt() time.Time {
	return s.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the unverified shop
func (s UnverifiedShop) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.UTC()
}

// Returns the name of the unverified shop
func (s UnverifiedShop) GetName() string {
	return s.Name
}

// Returns the phone number of the unverified shop
func (s UnverifiedShop) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the account balance in cents of the unverified shop
func (s UnverifiedShop) GetAccountBalanceInCents() int64 {
	return int64(s.AccountBalanceInCents)
}

// Returns the security password of the unverified shop
func (s UnverifiedShop) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the unverified shop
func (s UnverifiedShop) GetPinCode() string {
	return s.PinCode
}

// Returns the m-pesa number of the unverified shop
func (s UnverifiedShop) GetMpesaNumber() string {
	return s.MpesaNumber
}

type PhoneNumberExists struct {
	Verified   bool
	Unverified bool
}
