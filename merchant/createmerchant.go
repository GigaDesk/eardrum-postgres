package merchant

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/merchant"
	"gorm.io/gorm"
)

//create an unverified merchant record
func CreateMerchant(s merchant.NewMerchant, Db *gorm.DB) (merchant.Merchant, error) {
    //check if the phone number already exists
	phonenumberexists, err := CheckMerchantPhoneNumber(Db, s.GetPhoneNumber())

	if err!=nil{
		return nil, errors.New("error checking new merchant phonenumber existence")
	}

	if phonenumberexists.Unverified{
		return nil, errors.New("merchant phone number already exists but is unverified")
	}

	if phonenumberexists.Verified{
		return nil, errors.New("merchant phone number already exists")
	}

	//create unverified merchant data
	unverifiedmerchant := &UnverifiedMerchant{
		UserName:        s.GetUserName(),
		PhoneNumber: s.GetPhoneNumber(),
		Password:    s.GetPassword(),
		AccountBalanceInCents: 0,
		MpesaNumber: s.GetMpesaNumber(),
	}

	//create an unverified merchant record in the database and return if operation succeeds
	if err := Db.Create(unverifiedmerchant).Error; err != nil {
		return nil, err
	}

	return unverifiedmerchant, nil
}


