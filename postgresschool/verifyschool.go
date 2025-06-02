package postgresschool

import (
	"github.com/GigaDesk/eardrum-interfaces/school"
	"gorm.io/gorm"
)

//Transforms an unverified school record to a verified school record
func VerifySchool(phoneNumber string, Db *gorm.DB) (school.School, error) {
	//declare an unverifiedschool variable
	var unverifiedschool *UnverifiedSchool

	// Find the first unverified school that matches the input phone number from the unverified school table
	if err := Db.Where("phone_number = ?", phoneNumber).First(&unverifiedschool).Error; err != nil {
		return nil, err
	}

	// transform the unverified school model into school model and copy it
	school := &School{
		Name:        unverifiedschool.Name,
		PhoneNumber: unverifiedschool.PhoneNumber,
		Password:    unverifiedschool.Password,
		Badge:       unverifiedschool.Badge,
		Website:     unverifiedschool.Website,
	}

	// take the newly transformed and copied school data and transfer it into the official verified school table
	if err := Db.Create(school).Error; err != nil {
		return nil, err
	}

	// delete the unverified school from the unverified school table
	if err := Db.Delete(unverifiedschool).Error; err != nil {
		return nil, err
	}

	return school, nil
}
