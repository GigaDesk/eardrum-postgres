package postgresstudent

import (
	"github.com/GigaDesk/eardrum-interfaces/student"
	"gorm.io/gorm"
)

func GetStudentWithRegistrationNumber(Db *gorm.DB, RegistrationNumber string) (student.Student, error) {
	//declare a student variable
	var student *Student

	// Find the first student that matches the input registration number from the student table

	if err := Db.Where("registration_number = ?", RegistrationNumber).First(&student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func GetStudentWithId(Db *gorm.DB, Id int) (student.Student, error) {
	var student *Student
	//fetch the record to be updated from the database
	if err := Db.First(&student, Id).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func GetStudents(Db *gorm.DB) ([]student.Student, error) {

	var students []*Student

	if err := Db.Find(&students).Error; err != nil {
		return nil, err
	}

	var studentlist []student.Student

	for _, s := range students {
		studentlist = append(studentlist, s)
	}

	return studentlist, nil
}
