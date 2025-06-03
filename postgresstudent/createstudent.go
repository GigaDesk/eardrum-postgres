package postgresstudent

import (
	"github.com/GigaDesk/eardrum-interfaces/student"
	"gorm.io/gorm"
)

// creates a student record
func CreateStudent(s student.NewStudent, Db *gorm.DB) (student.Student, error) {
	//create student data

	student := &Student{
		Name:               s.GetName(),
		PhoneNumber:        s.GetPhoneNumber(),
		RegistrationNumber: s.GetRegistrationNumber(),
		Password:           s.GetPassword(),
		DateOfAdmission:    s.GetDateOfAdmission(),
		DateOfBirth:        s.GetDateofBirth(),
		ProfilePicture:     s.GetProfilePicture(),
		AccountBalanceInCents: 0,
	}

	//create a student record in the database and return if operation succeeds
	if err := Db.Create(student).Error; err != nil {
		return nil, err
	}

	return student, nil
}
