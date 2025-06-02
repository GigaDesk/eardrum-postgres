package mockschool

import "time"

var SchoolNode = MockSchool{
	Id:          1,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Kisii School",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
	Badge:       "badge",
	Website:     "Kisiischool@ac.ke",
}

var SamePhoneNumberSchoolNode = MockSchool{
	Id:          2,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Kisii School",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
	Badge:       "badge",
	Website:     "Kisiischool@ac.ke",
}

var SameIdSchoolNode = MockSchool{
	Id:          1,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Kisii School",
	PhoneNumber: "+254719226154",
	Password:    "kisbhdbcvukbqiyde327&",
	Badge:       "badge",
	Website:     "Kisiischool@ac.ke",
}

var MultipleSchoolNodes = []MockSchool{
	{
	Id:          3,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Kisii School",
	PhoneNumber: "+254719226155",
	Password:    "kisbhdbcvukbqiyde327&",
	Badge:       "badge",
	Website:     "Kisiischool@ac.ke",
	},
	{
		Id:          4,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Kisii School",
		PhoneNumber: "+254719226156",
		Password:    "kisbhdbcvukbqiyde327&",
		Badge:       "badge",
		Website:     "Kisiischool@ac.ke",
	},
	{
		Id:          5,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Kisii School",
		PhoneNumber: "+254719226157",
		Password:    "kisbhdbcvukbqiyde327&",
		Badge:       "badge",
		Website:     "Kisiischool@ac.ke",
	},
	{
		Id:          6,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "London School of Business",
		PhoneNumber: "+254719226158",
		Password:    "kisbhdbcvukbqiyde327&",
		Badge:       "badge",
		Website:     "Kisiischool@ac.ke",
	},
}

var UpdatedSchool = MockSchool{
	Id:          4,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "London Business School",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
	Badge:       "badge",
	Website:     "londonschool@ac.ke",
}