package postgresproduct

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name                  string    // Name of the product
	PricePerUnitInCents int64     // The product's price per unit in cents
}

// Returns the unique ID of the product
func (p Product) GetID() int64 {
	return int64(p.ID)
}

// Returns the creation timestamp of the product
func (p Product) GetCreatedAt() time.Time {
	return p.CreatedAt.UTC()
}

// Returns the update timestamp of the product
func (p Product) GetUpdatedAt() time.Time {
	return p.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the product
func (p Product) GetDeletedAt() time.Time {
	return p.DeletedAt.Time.UTC()
}

// Returns the name of the product
func (p Product) GetName() string {
	return p.Name
}

// Returns the price per unit in cents of the product
func (p Product) GetPricePerUnitInCents() int64 {
	return p.PricePerUnitInCents
}