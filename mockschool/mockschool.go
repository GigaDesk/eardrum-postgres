package mockschool

import "time"

type MockSchool struct {
	Id          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	PhoneNumber string
	Password    string
	Badge       string
	Website     string
}




func (m MockSchool) GetID() int64 {
	return int64(m.Id)
}
func (m MockSchool) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockSchool) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockSchool) GetName() string {
	return m.Name
}
func (m MockSchool) GetPhoneNumber() string {
	return m.PhoneNumber
}
func (m MockSchool) GetPassword() string {
	return m.Password
}
func (m MockSchool) GetBadge() string {
	return m.Badge
}
func (m MockSchool) GetWebsite() string {
	return m.Website
}
func (m MockSchool) GetDeletedAt() time.Time {
	return time.Time{}
}