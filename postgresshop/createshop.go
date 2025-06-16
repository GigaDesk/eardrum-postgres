package postgresshop

import (
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

//create a shop record 
func CreateShop(s shop.NewShop, Db *gorm.DB) (shop.Shop, error) {
	//create shop data
	shop := &Shop{
		Name:        s.GetName(),
		PhoneNumber: s.GetPhoneNumber(),
		Password:    s.GetPassword(),
		AccountBalanceInCents: 0,
	}

	//create a shop record in the database and return if operation succeeds
	if err := Db.Create(shop).Error; err != nil {
		return nil, err
	}

	return shop, nil
}


