package postgresschool

import (
	"github.com/GigaDesk/eardrum-interfaces/school"
	"gorm.io/gorm"
)

func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (school.School, error) {
	var school *School
	//fetch the record to be updated from the database
	if err := Db.First(&school, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&school).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&school, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return school, nil
}
