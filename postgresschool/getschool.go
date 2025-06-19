package postgresschool

import (
	"github.com/GigaDesk/eardrum-interfaces/school"
	"gorm.io/gorm"
)

func GetSchoolWithPhoneNumber(Db *gorm.DB, phoneNumber string) (school.School, error) {
	//declare a school variable
	var school *School
	// Find the first school that matches the phone number from the school table

	if err := Db.Where("phone_number = ?", phoneNumber).First(&school).Error; err != nil {
		return nil, err
	}

	return school, nil
}

func GetSchoolWithId(Db *gorm.DB, Id int) (school.School, error) {
	var school *School
	//fetch the record to be updated from the database
	if err := Db.First(&school, Id).Error; err != nil {
		return nil, err
	}

	return school, nil
}

func GetSchools(Db *gorm.DB) ([]school.School, error) {

	var schools []*School

	if err := Db.Find(&schools).Error; err != nil {
		return nil, err
	}

	var schoollist []school.School

	for _, s := range schools {
		schoollist = append(schoollist, s)
	}

	return schoollist, nil
}
