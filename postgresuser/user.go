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
func (u User) GetID() int64 {
	return int64(u.ID)
}

// Returns the creation timestamp of the user
func (u User) GetCreatedAt() time.Time {
	return u.CreatedAt.UTC()
}

// Returns the update timestamp of the user
func (u User) GetUpdatedAt() time.Time {
	return u.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the user
func (u User) GetDeletedAt() time.Time {
	return u.DeletedAt.Time.UTC()
}

// Returns the name of the user
func (u User) GetName() string {
	return u.Name
}

// Returns the phone number of the user
func (u User) GetPhoneNumber() string {
	return u.PhoneNumber
}

// Returns the account balance in cents of the user
func (u User) GetAccountBalanceInCents() int64 {
	return u.AccountBalanceInCents
}

// Returns the security password of the user
func (u User) GetPassword() string {
	return u.Password
}

// Returns the security PIN code of the user
func (u User) GetPinCode() string {
	return u.PinCode
}


type UnverifiedUser struct {
	gorm.Model
	Name                  string    // Name of the unverified user
	PhoneNumber           string    // Phone number of the unverified user
	Password              string    // Security password of the unverified user
	AccountBalanceInCents int64     // The unverified user's account balance in cents
	PinCode               string    // The security PIN code of the unverified user
}

// Returns the unique ID of the unverified user
func (u UnverifiedUser) GetID() int64 {
	return int64(u.ID)
}

// Returns the creation timestamp of the unverified user
func (u UnverifiedUser) GetCreatedAt() time.Time {
	return u.CreatedAt.UTC()
}

// Returns the update timestamp of the unverified user
func (u UnverifiedUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the unverified user
func (u UnverifiedUser) GetDeletedAt() time.Time {
	return u.DeletedAt.Time.UTC()
}

// Returns the name of the unverified user
func (u UnverifiedUser) GetName() string {
	return u.Name
}

// Returns the phone number of the unverified user
func (u UnverifiedUser) GetPhoneNumber() string {
	return u.PhoneNumber
}

// Returns the account balance in cents of the unverified user
func (u UnverifiedUser) GetAccountBalanceInCents() int64 {
	return u.AccountBalanceInCents
}

// Returns the security password of the unverified user
func (u UnverifiedUser) GetPassword() string {
	return u.Password
}

// Returns the security PIN code of the unverified user
func (u UnverifiedUser) GetPinCode() string {
	return u.PinCode
}

type PhoneNumberExists struct {
	Verified bool
	Unverified bool
}