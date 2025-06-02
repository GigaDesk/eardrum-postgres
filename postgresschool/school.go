package postgresschool

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name        string // Name of the school
	PhoneNumber string // Phone number of the school
	Badge       string // Badge or identifier associated with the school
	Website     string // Website URL of the school
	Password    string // Security password of the school
}

// Returns the unique ID of the school
func (s School) GetID() int64 {
	return int64(s.ID)
}

// Returns the creation timestamp of the school
func (s School) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// Returns the update timestamp of the school
func (s School) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// Returns the deletion timestamp of the school
func (s School) GetDeletedAt() time.Time {
	return s.DeletedAt.Time
}

// Returns the name of the school
func (s School) GetName() string {
	return s.Name
}

// Returns the phone number of the school
func (s School) GetPhoneNumber() string {
	return s.PhoneNumber
}

// Returns the badge url of the school
func (s School) GetBadge() string {
	return s.Badge
}

// Returns the website url of the school
func (s School) GetWebsite() string {
	return s.Website
}

// Returns the security password of the school
func (s School) GetPassword() string {
	return s.Password
}

type UnverifiedSchool struct {
	gorm.Model
	Name        string // Name of the school
	PhoneNumber string // Phone number of the school
	Badge       string // Badge or identifier associated with the school
	Website     string // Website URL of the school
	Password    string // Security password of the school
}

// Returns the unique ID of the unverified school
func (u UnverifiedSchool) GetID() int64 {
	return int64(u.ID)
}

// Returns the creation timestamp of the unverified school
func (u UnverifiedSchool) GetCreatedAt() time.Time {
	return u.CreatedAt
}

// Returns the update timestamp of the unverified school
func (u UnverifiedSchool) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

// Returns the deletion timestamp of the unverified school
func (u UnverifiedSchool) GetDeletedAt() time.Time {
	return u.DeletedAt.Time
}

// Returns the name of the unverified school
func (u UnverifiedSchool) GetName() string {
	return u.Name
}

// Returns the phone number of the unverified school
func (u UnverifiedSchool) GetPhoneNumber() string {
	return u.PhoneNumber
}

// Returns the badge url of the unverified school
func (u UnverifiedSchool) GetBadge() string {
	return u.Badge
}

// Returns the website url of the unverified school
func (u UnverifiedSchool) GetWebsite() string {
	return u.Website
}

// Returns the security password of the unverified school
func (u UnverifiedSchool) GetPassword() string {
	return u.Password
}

type PhoneNumberExists struct {
  Verified bool
  Unverified bool
}