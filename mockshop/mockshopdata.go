package mockshop

import (
	"time"
)

var MultipleShopNodes = []MockShop{
	{
	Id:          3,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-025/2020",
	Name:        "Canteen A",
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
		Name:        "Canteen B",
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
		Name:        "Canteen C",
		PhoneNumber: "+254719226157",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "Kisiischool@ac.ke",
	},
	{
		Id:          6,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		RegistrationNumber: "SCT-251-025/2023",
		Name:        "Canteen D",
		PhoneNumber: "+254719226158",
		Password:    "kisbhdbcvukbqiyde327&",
		DateOfAdmission: time.Now().UTC(),
		DateofBirth: time.Now().UTC(),
		ProfilePicture:     "Kisiischool@ac.ke",
	},
}

var UpdatedShop = MockShop{
	Id:          6,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	RegistrationNumber: "SCT-251-025/2023",
	Name:        "Book Shop",
	PhoneNumber: "+254116063319",
	Password:    "kisbhdbcvukbqiyde327&",
	DateOfAdmission: time.Now().UTC(),
	DateofBirth: time.Now().UTC(),
	ProfilePicture:     "Kisiischool@ac.ke",
	Account_balance_in_cents: 8000,
}