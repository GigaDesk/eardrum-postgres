package postgresuser

import (
	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

// Transforms an unverified user record to a verified user record
func VerifyUser(phoneNumber string, Db *gorm.DB) (user.User, error) {
	//declare an unverified user variable
	var unverifieduser *UnverifiedUser

	// Find the first unverified user that matches the input phone number from the unverified user table
	if err := Db.Where("phone_number = ?", phoneNumber).First(&unverifieduser).Error; err != nil {
		return nil, err
	}

	// transform the unverified user model into user model and copy it
	user := &User{
		Name:                  unverifieduser.Name,
		PhoneNumber:           unverifieduser.PhoneNumber,
		Password:              unverifieduser.Password,
		AccountBalanceInCents: unverifieduser.AccountBalanceInCents,
		PinCode:               unverifieduser.PinCode,
	}

	// take the newly transformed and copied user data and transfer it into the official verified user table
	if err := Db.Create(user).Error; err != nil {
		return nil, err
	}

	// delete the unverified user from the unverified user table
	if err := Db.Delete(unverifieduser).Error; err != nil {
		return nil, err
	}

	return user, nil
}
