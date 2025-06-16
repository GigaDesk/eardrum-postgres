package postgresuser

import (
	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

// creates an unverified user record
func CreateUser(s user.NewUser, Db *gorm.DB) (user.User, error) {
	//create an unverified user data

	unverifieduser := &UnverifiedUser{
		Name:               s.GetName(),
		PhoneNumber:        s.GetPhoneNumber(),
		Password:           s.GetPassword(),
		AccountBalanceInCents: 0,
	}

	//create an unverified user record in the database and return if operation succeeds
	if err := Db.Create(unverifieduser).Error; err != nil {
		return nil, err
	}

	return unverifieduser, nil
}
