package mocktransaction

import "time"


type MockTransaction struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	TotalAmountInCents int64
	BalanceBeforeInCents int64
	BalanceAfterInCents  int64
}




func (m MockTransaction) GetID() int64 {
	return m.Id
}
func (m MockTransaction) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockTransaction) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockTransaction) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockTransaction) GetTotalAmountInCents() int64 {
	return m.TotalAmountInCents
}
func (m MockTransaction) GetBalanceBeforeInCents() int64 {
	return m.BalanceBeforeInCents
}
func (m MockTransaction) GetBalanceAfterInCents() int64 {
	return m.BalanceAfterInCents
}