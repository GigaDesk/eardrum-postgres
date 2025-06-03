package mockstudent

import "time"


type MockStudent struct {
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




func (m MockStudent) GetID() int64 {
	return m.Id
}
func (m MockStudent) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockStudent) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockStudent) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockStudent) GetRegistrationNumber() string {
	return m.RegistrationNumber
}
func (m MockStudent) GetName() string {
	return m.Name
}
func (m MockStudent) GetPhoneNumber() string {
	return m.PhoneNumber
}
func (m MockStudent) GetPassword() string {
	return m.Password
}
func (m MockStudent) GetDateOfAdmission() time.Time {
	return m.DateOfAdmission
}
func (m MockStudent) GetDateofBirth() time.Time {
	return m.DateofBirth
}
func (m MockStudent) GetProfilePicture() string {
	return m.ProfilePicture
}
func (m MockStudent) GetAccountBalanceInCents() int64{
    return m.Account_balance_in_cents
}
func (m MockStudent) GetPinCode() string{
    return ""
}
