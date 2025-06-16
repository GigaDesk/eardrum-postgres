package postgresshop

import (
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

//create an unverified shop record 
func CreateShop(s shop.NewShop, Db *gorm.DB) (shop.Shop, error) {
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


