package postgresshop

import (
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

func GetShopWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (shop.Shop, error) {
	//declare a shop variable
	var shop *Shop

	// Find the first shop that matches the input phonenumber from the shop table

	if err := Db.Where("phone_number = ?", PhoneNumber).First(&shop).Error; err != nil {
		return nil, err
	}

	return shop, nil
}

//Gets a shop by its unique id
func GetShopWithId(Db *gorm.DB, Id int) (shop.Shop, error) {
	var shop *Shop
	//fetch the record to be updated from the database
	if err := Db.First(&shop, Id).Error; err != nil {
		return nil, err
	}

	return shop, nil
}

//Gets all the shops registered in the database
func GetShops(Db *gorm.DB) ([]shop.Shop, error) {

	var shops []*Shop

	if err := Db.Find(&shops).Error; err != nil {
		return nil, err
	}

	var shoplist []shop.Shop

	for _, s := range shops {
		shoplist = append(shoplist, s)
	}

	return shoplist, nil
}
