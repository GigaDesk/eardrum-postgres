package postgresstudent

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name                  string    // Name of the student
	PhoneNumber           string    // Phone number of the student
	RegistrationNumber    string    // Registration number of the student
	DateOfAdmission       time.Time // The student's date of admission
	DateOfBirth           time.Time // The student's date of birth
	ProfilePicture        string    // The student's profile picture
	Password              string    // Security password of the student
	AccountBalanceInCents int64     // The student's account balance in cents
	PinCode               string    // The security PIN code of the student
}

// Returns the unique ID of the student
func (s Student) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the student
func (s Student) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// Returns the update timestamp of the student
func (s Student) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// Returns the deletion timestamp of the student
func (s Student) GetDeletedAt() time.Time {
	return s.DeletedAt.Time
}

// Returns the name of the student
func (s Student) GetName() string {
	return s.Name
}

// Returns the phone number of the student
func (s Student) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the account balance in cents of the student
func (s Student) GetAccountBalanceInCents() int64 {
	return s.AccountBalanceInCents
}

// Returns the date of birth of the student
func (s Student) GetDateofBirth() time.Time {
	return s.DateOfBirth
}

// Returns the date of admission of the student
func (s Student) GetDateOfAdmission() time.Time {
	return s.DateOfAdmission
}

// Returns the security password of the student
func (s Student) GetPassword() string {
	return s.Password
}

// Returns the security PIN code of the student
func (s Student) GetPinCode() string {
	return s.PinCode
}

// Returns the profile picture url of the student
func (s Student) GetProfilePicture() string {
	return s.ProfilePicture
}

// Returns the profile picture url of the student
func (s Student) GetRegistrationNumber() string {
	return s.RegistrationNumber
}
