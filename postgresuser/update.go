package postgresuser

import (
	"github.com/GigaDesk/eardrum-interfaces/user"
	"gorm.io/gorm"
)

func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (user.User, error) {
	var user *User
	//fetch the record to be updated from the database
	if err := Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&user).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return user, nil
}


func UpdatePinCode(Db *gorm.DB, encryptedpincode string, id int) (user.User, error) {
	var user *User
	//fetch the record to be updated from the database
	if err := Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&user).Updates(map[string]interface{}{"pin_code": encryptedpincode}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return user, nil
}