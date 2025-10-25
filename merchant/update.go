package merchant

import (
	"github.com/GigaDesk/eardrum-interfaces/merchant"
	"gorm.io/gorm"
)

func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (merchant.Merchant, error) {
	var merchant *Merchant
	//fetch the record to be updated from the database
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&merchant).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return merchant, nil
}

func UpdatePinCode(Db *gorm.DB, encryptedpincode string, id int) (merchant.Merchant, error) {
	var merchant *Merchant
	//fetch the record to be updated from the database
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&merchant).Updates(map[string]interface{}{"pin_code": encryptedpincode}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return merchant, nil
}

func UpdateMpesaNumber(Db *gorm.DB, new_mpesa_number string, id int) (merchant.Merchant, error) {
	var merchant *Merchant
	//fetch the record to be updated from the database
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&merchant).Updates(map[string]interface{}{"mpesa_number ": new_mpesa_number}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return merchant, nil
}
