package postgresschool

import (
	"github.com/GigaDesk/eardrum-interfaces/school"
	"gorm.io/gorm"
)

//create an unverified school record 
func CreateSchool(s school.NewSchool, Db *gorm.DB) (school.School, error) {
	//create an unverified school data
	unverifiedschool := &UnverifiedSchool{
		Name:        s.GetName(),
		PhoneNumber: s.GetPhoneNumber(),
		Password:    s.GetPassword(),
		Badge:       s.GetBadge(),
		Website:     s.GetWebsite(),
	}

	//create an unverified school record in the database and return if operation succeeds
	if err := Db.Create(unverifiedschool).Error; err != nil {
		return nil, err
	}

	return unverifiedschool, nil
}


