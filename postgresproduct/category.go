package postgresproduct

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string // Name of the  product category
	Description string // The category's description
}

// Returns the unique ID of the product category
func (c Category) GetID() int64 {
	return int64(c.ID)
}

// Returns the creation timestamp of the product category
func (c Category) GetCreatedAt() time.Time {
	return c.CreatedAt.UTC()
}

// Returns the update timestamp of the product category
func (c Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the product category
func (c Category) GetDeletedAt() time.Time {
	return c.DeletedAt.Time.UTC()
}

// Returns the name of the product category
func (c Category) GetName() string {
	return c.Name
}

// Returns the description of the product category
func (c Category) GetDescription() string {
	return c.Description
}
