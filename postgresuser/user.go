package postgresuser

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                  string    // Name of the user
	PhoneNumber           string    // Phone number of the user
	Password              string    // Security password of the user
	AccountBalanceInCents int64     // The user's account balance in cents
	PinCode               string    // The security PIN code of the user
}

// Returns the unique ID of the user
func (s User) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the user
func (s User) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// Returns the update timestamp of the user
func (s User) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// Returns the deletion timestamp of the user
func (s User) GetDeletedAt() time.Time {
	return s.DeletedAt.Time
}

// Returns the name of the user
func (s User) GetName() string {
	return s.Name
}

// Returns the phone number of the user
func (s User) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the account balance in cents of the user
func (s User) GetAccountBalanceInCents() int64 {
	return s.AccountBalanceInCents
}

// Returns the security password of the user
func (s User) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the user
func (s User) GetPinCode() string {
	return s.PinCode
}
