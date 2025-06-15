package postgresuser

import (
	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

// creates a user record
func CreateUser(s user.NewUser, Db *gorm.DB) (user.User, error) {
	//create user data

	user := &User{
		Name:               s.GetName(),
		PhoneNumber:        s.GetPhoneNumber(),
		Password:           s.GetPassword(),
		AccountBalanceInCents: 0,
	}

	//create a user record in the database and return if operation succeeds
	if err := Db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
