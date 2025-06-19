package mockshop

import "time"


type MockShop struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	RegistrationNumber string
	Name string
	PhoneNumber string
	Password string
	DateOfAdmission time.Time
	DateofBirth time.Time
	ProfilePicture string
	Account_balance_in_cents int64
}




func (m MockShop) GetID() int64 {
	return m.Id
}
func (m MockShop) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockShop) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockShop) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockShop) GetRegistrationNumber() string {
	return m.RegistrationNumber
}
func (m MockShop) GetName() string {
	return m.Name
}
func (m MockShop) GetPhoneNumber() string {
	return m.PhoneNumber
}
func (m MockShop) GetPassword() string {
	return m.Password
}
func (m MockShop) GetDateOfAdmission() time.Time {
	return m.DateOfAdmission
}
func (m MockShop) GetDateofBirth() time.Time {
	return m.DateofBirth
}
func (m MockShop) GetProfilePicture() string {
	return m.ProfilePicture
}
func (m MockShop) GetAccountBalanceInCents() int64{
    return m.Account_balance_in_cents
}
func (m MockShop) GetPinCode() string{
    return ""
}