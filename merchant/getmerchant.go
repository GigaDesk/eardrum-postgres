package merchant

import (
	"github.com/GigaDesk/eardrum-interfaces/merchant"
	"gorm.io/gorm"
)

func GetMerchantWithPhoneNumber(Db *gorm.DB, PhoneNumber string) (merchant.Merchant, error) {
	//declare a merchant variable
	var merchant *Merchant

	// Find the first merchant that matches the input phonenumber from the merchant table

	if err := Db.Where("phone_number = ?", PhoneNumber).First(&merchant).Error; err != nil {
		return nil, err
	}

	return merchant, nil
}

//Gets a merchant by its unique id
func GetMerchantWithId(Db *gorm.DB, Id int) (merchant.Merchant, error) {
	var merchant *Merchant
	//fetch the record to be updated from the database
	if err := Db.First(&merchant, Id).Error; err != nil {
		return nil, err
	}

	return merchant, nil
}

//Gets all the merchants registered in the database
func GetMerchants(Db *gorm.DB) ([]merchant.Merchant, error) {

	var merchants []*Merchant

	if err := Db.Find(&merchants).Error; err != nil {
		return nil, err
	}

	var merchantlist []merchant.Merchant

	for _, s := range merchants {
		merchantlist = append(merchantlist, s)
	}

	return merchantlist, nil
}
