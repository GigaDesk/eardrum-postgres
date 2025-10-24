package user

import (
	"errors" // Standard Go errors package
	"gorm.io/gorm"
	"github.com/GigaDesk/eardrum-interfaces/user"
)

// GetUserWithPhoneNumber finds a user by phone number.
func GetUserWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (user.User, error) {
	var u *User

	// Find the first user that matches the input phonenumber
	if err := Db.Where("phone_number = ?", PhoneNumber).First(&u).Error; err != nil {
		
		// 1. Check if the error is the known "Not Found" condition
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound("phone_number", PhoneNumber) // Returns 404
		}

		// 2. All other errors (connection, query syntax, etc.) -> 500 Internal
		return nil, ErrDBLookupFailure("Failed to execute query for user phone number.", err)
	}

	return u, nil
}

// GetUserWithId finds a user by ID.
func GetUserWithId(Db *gorm.DB, Id int) (user.User, error) {
	var u *User
	
	// Fetch the record by primary key
	if err := Db.First(&u, Id).Error; err != nil {
		
		// 1. Check for Not Found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound("ID", Id) // Returns 404
		}

		// 2. All other errors -> 500 Internal
		return nil, ErrDBLookupFailure("Failed to execute query for user ID.", err)
	}

	return u, nil
}

// GetUsers returns all users.
func GetUsers(Db *gorm.DB) ([]user.User, error) {
	var users []*User

	// Find all records
	if err := Db.Find(&users).Error; err != nil {
		// Note: Db.Find does NOT return ErrRecordNotFound if the table is empty.
		// It only returns an error on connection or query issue.
		return nil, ErrDBLookupFailure("Failed to retrieve list of all users.", err)
	}

	// Transform [](*User) to []user.User (assuming *User implements user.User)
	var userslist []user.User
	for _, u := range users {
		userslist = append(userslist, u)
	}

	return userslist, nil
}