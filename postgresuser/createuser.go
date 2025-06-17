package postgresuser

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

// creates an unverified user record
func CreateUser(s user.NewUser, Db *gorm.DB) (user.User, error) {
	//check if the phone number already exists
	phonenumberexists, err := CheckUserPhoneNumber(Db, s.GetPhoneNumber())

	if err!=nil{
		return nil, errors.New("error checking new user phonenumber existence")
	}

	if phonenumberexists.Unverified{
		return nil, errors.New("user phone number already exists but is unverified")
	}

	if phonenumberexists.Verified{
		return nil, errors.New("user phone number already exists")
	}
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
