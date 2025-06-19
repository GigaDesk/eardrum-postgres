package postgresuser

import (
	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

func GetUserWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (user.User, error) {
	//declare a user variable
	var user *User

	// Find the first user that matches the input phonenumber from the user table

	if err := Db.Where("phone_number = ?", PhoneNumber).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserWithId(Db *gorm.DB, Id int) (user.User, error) {
	var user *User
	//fetch the record to be updated from the database
	if err := Db.First(&user, Id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers(Db *gorm.DB) ([]user.User, error) {

	var users []*User

	if err := Db.Find(&users).Error; err != nil {
		return nil, err
	}

	var userslist []user.User

	for _, s := range users {
		userslist = append(userslist, s)
	}

	return userslist, nil
}
