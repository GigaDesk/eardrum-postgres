package mockproduct

import (
	"time"

	"github.com/GigaDesk/eardrum-interfaces/product"
)

var p product.Category

type MockCategory struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name string
	Description string
}


func (m MockCategory) GetID() int64 {
	return m.Id
}
func (m MockCategory) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockCategory) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockCategory) GetName() string {
	return m.Name
}
func (m MockCategory) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockCategory) GetDescription() string{
    return m.Description
}
