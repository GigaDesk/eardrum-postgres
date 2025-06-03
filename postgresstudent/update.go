package postgresstudent

import (
	"github.com/GigaDesk/eardrum-interfaces/student"
	"gorm.io/gorm"
)

func UpdatePassword(Db *gorm.DB, encryptedpassword string, id int) (student.Student, error) {
	var student *Student
	//fetch the record to be updated from the database
	if err := Db.First(&student, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&student).Updates(map[string]interface{}{"password": encryptedpassword}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&student, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return student, nil
}


func UpdatePinCode(Db *gorm.DB, encryptedpincode string, id int) (student.Student, error) {
	var student *Student
	//fetch the record to be updated from the database
	if err := Db.First(&student, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&student).Updates(map[string]interface{}{"pin_code": encryptedpincode}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&student, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return student, nil
}