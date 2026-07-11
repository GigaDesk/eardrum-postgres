package user

import (
	pgerror "errors" // Standard Go errors package
	"gorm.io/gorm"
	"github.com/GigaDesk/eardrum-interfaces/user"
	"github.com/GigaDesk/eardrum-interfaces/errors"
)

// GetUserWithPhoneNumber finds a user by phone number.
func GetUserWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (user.User, error) {
	var u *User

	// Find the first user that matches the input phonenumber
	if err := Db.Where("phone_number = ?", PhoneNumber).First(&u).Error; err != nil {
		
			// 1. Check for the known "Not Found" condition
		if pgerror.Is(err, gorm.ErrRecordNotFound) {
			// Returns 404 Not Found
			err1 := errors.New(errors.EARUserNotFoundByPhone, err)
			err1.Log()
			return nil, err1
		}

		// 2. All other errors (connection, query syntax, etc.) -> 500 Internal
		err1 := errors.New(errors.EARUserLookupFailedByPhone, err)
		err1.Log()
		return nil, err1
	}

	return u, nil
}

// GetUserWithUsername finds a user by username.
func GetUserWithUsername(Db *gorm.DB, Username string) (user.User, error) {
    var u *User

    // Find the first user that matches the input username
    if err := Db.Where("user_name = ?", Username).First(&u).Error; err != nil {
        
        // 1. Check for the known "Not Found" condition
        if pgerror.Is(err, gorm.ErrRecordNotFound) {
            // Returns 404 Not Found (Assuming EARUserNotFoundByUsername exists in your interfaces)
            err1 := errors.New(errors.EARUserNotFoundByUsername, err)
            err1.Log()
            return nil, err1
        }

        // 2. All other errors (connection, query syntax, etc.) -> 500 Internal
        err1 := errors.New(errors.EARUserLookupFailedByUsername, err)
        err1.Log()
        return nil, err1
    }

    return u, nil
}
// GetUsers retrieves a paginated list of users registered in the database.
// - limit: The maximum number of records to return (e.g., 5)
// - offset: The number of records to skip before starting to return (e.g., 0 for the first page)
func GetUsers(Db *gorm.DB, limit int, offset int) ([]user.User, error) {
    var users []*User

    // Find all records with limit and offset applied
    if err := Db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
        // Db.Find only returns an error on connection or query issue, not if the table is empty.
        err1 := errors.New(errors.EARUserListRetrievalFailed, err)
        err1.Log()
        return nil, err1
    }

    // Transform [](*User) to []user.User
    var userslist []user.User
    for _, u := range users {
        userslist = append(userslist, u)
    }

    return userslist, nil
}


