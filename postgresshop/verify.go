package postgresshop

import (
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

// Transforms an unverified shop record to a verified shop record
func VerifyShop(phoneNumber string, Db *gorm.DB) (shop.Shop, error) {
	//declare an unverified shop variable
	var unverifiedshop *UnverifiedShop

	// Find the first unverified shop that matches the input phone number from the unverified shop table
	if err := Db.Where("phone_number = ?", phoneNumber).First(&unverifiedshop).Error; err != nil {
		return nil, err
	}

	// transform the unverified shop model into shop model and copy it
	shop := &Shop{
		Name:                  unverifiedshop.Name,
		PhoneNumber:           unverifiedshop.PhoneNumber,
		Password:              unverifiedshop.Password,
		AccountBalanceInCents: unverifiedshop.AccountBalanceInCents,
		PinCode:               unverifiedshop.PinCode,
	}

	// take the newly transformed and copied shop data and transfer it into the official verified shop table
	if err := Db.Create(shop).Error; err != nil {
		return nil, err
	}

	// delete the unverified shop from the unverified shop table
	if err := Db.Delete(unverifiedshop).Error; err != nil {
		return nil, err
	}

	return shop, nil
}