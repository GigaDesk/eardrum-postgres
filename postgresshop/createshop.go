package postgresshop

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

//create an unverified shop record
func CreateShop(s shop.NewShop, Db *gorm.DB) (shop.Shop, error) {
    //check if the phone number already exists
	phonenumberexists, err := CheckShopPhoneNumber(Db, s.GetPhoneNumber())

	if err!=nil{
		return nil, errors.New("error checking new shop phonenumber existence")
	}

	if phonenumberexists.Unverified{
		return nil, errors.New("shop phone number already exists but is unverified")
	}

	if phonenumberexists.Unverified{
		return nil, errors.New("shop phone number already exists")
	}

	//create unverified shop data
	unverifiedshop := &UnverifiedShop{
		Name:        s.GetName(),
		PhoneNumber: s.GetPhoneNumber(),
		Password:    s.GetPassword(),
		AccountBalanceInCents: 0,
	}

	//create an unverified shop record in the database and return if operation succeeds
	if err := Db.Create(unverifiedshop).Error; err != nil {
		return nil, err
	}

	return unverifiedshop, nil
}


