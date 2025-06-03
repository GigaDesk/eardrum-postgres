package postgresshop

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name                  string    // Name of the shop
	PhoneNumber           string    // Phone number of the shop
	Password              string    // Security password of the shop
	AccountBalanceInCents int64     // The shop's account balance in cents
	PinCode               string    // The security PIN code of the shop
}

// Returns the unique ID of the shop
func (s Shop) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the shop
func (s Shop) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// Returns the update timestamp of the shop
func (s Shop) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// Returns the deletion timestamp of the shop
func (s Shop) GetDeletedAt() time.Time {
	return s.DeletedAt.Time
}

// Returns the name of the shop
func (s Shop) GetName() string {
	return s.Name
}

// Returns the phone number of the shop
func (s Shop) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the account balance in cents of the shop
func (s Shop) GetAccountBalanceInCents() int64 {
	return s.AccountBalanceInCents
}

// Returns the security password of the shop
func (s Shop) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the shop
func (s Shop) GetPinCode() string {
	return s.PinCode
}
