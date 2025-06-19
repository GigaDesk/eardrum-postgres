package mockstudent

import "time"

//Mock data for carrying out student tests

var StudentNode = MockStudent{
	Id:          1,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-025/2019",
	Name:        "Leon Kenyaga",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
	DateOfAdmission: time.Now().UTC(),
	DateofBirth: time.Now().UTC(),
	ProfilePicture: "https://image.url",
}

var SameRegistrationNumberStudentNode = MockStudent{
	Id:          2,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-025/2019",
	Name:        "Kisii School",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
	DateOfAdmission: time.Now().UTC(),
	DateofBirth: time.Now().UTC(),
	ProfilePicture:     "Kisiischool@ac.ke",
}

var SameIdStudentNode = MockStudent{
	Id:          1,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-035/2019",
	Name:        "Kisii School",
	PhoneNumber: "+254719226154",
	Password:    "kisbhdbcvukbqiyde327&",
	DateOfAdmission: time.Now().UTC(),
	DateofBirth: time.Now().UTC(),
	ProfilePicture:     "Kisiischool@ac.ke",
}

var MultipleStudentNodes = []MockStudent{
	{
	Id:          3,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-025/2020",
	Name:        "Arnold Osoro",
	PhoneNumber: "+254719226155",
	Password:    "kisbhdbcvukbqiyde327&",
	DateOfAdmission: time.Now().UTC(),
	DateofBirth: time.Now().UTC(),
	ProfilePicture:     "Kisiischool@ac.ke",
	},
	{
		Id:          4,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		RegistrationNumber: "SCT-251-025/2021",
		Name:        "Julius Nyakweba",
		PhoneNumber: "+254719226156",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "Kisiischool@ac.ke",
	},
	{
		Id:          5,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		RegistrationNumber: "SCT-251-025/2022",
		Name:        "Carlos Okumu",
		PhoneNumber: "+254719226157",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "Kisiischool@ac.ke",
		Account_balance_in_cents: 500,
	},
	{
		Id:          6,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		RegistrationNumber: "SCT-251-025/2023",
		Name:        "Greg Otieno",
		PhoneNumber: "+254719226158",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "Kisiischool@ac.ke",
	},
}

var UpdatedStudent = MockStudent{
		Id:          5,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		RegistrationNumber: "SCT-251-025/2072",
		Name:        "Carlos Okumu Nyambati",
		PhoneNumber: "+254705136690",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "carlos@ac.ke",
		Account_balance_in_cents: 5000,
}