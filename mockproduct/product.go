package mockproduct

import (
	"time"
)


type MockProduct struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name string
	Price_per_unit_in_cents int64
}


func (m MockProduct) GetID() int64 {
	return m.Id
}
func (m MockProduct) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockProduct) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockProduct) GetName() string {
	return m.Name
}
func (m MockProduct) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockProduct) GetPricePerUnitInCents() int64{
    return m.Price_per_unit_in_cents
}
