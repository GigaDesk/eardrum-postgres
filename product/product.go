package product

import (
	"errors"
	"time"

	"github.com/GigaDesk/eardrum-postgres/merchant"
	"gorm.io/gorm"
)

// Product represents a product that a shop sells.
type Product struct {
	gorm.Model
	Name                string    `gorm:"not null;uniqueIndex:idx_product_shop"` // The name of the product. Must be unique within a single shop.
	PricePerUnitInCents uint      `gorm:"not null"`                              // The product's price per unit. A uint cannot be negative.
	MerchantID              uint      `gorm:"not null;uniqueIndex:idx_product_shop"` // The foreign key linking this product to its parent shop.
	Merchant                merchant.Merchant `gorm:"foreignKey:MerchantID"`                     // GORM association to the Merchant model.
	CategoryID          *uint     // The foreign key linking this product to its category. A product can be uncategorized.
	Category            *Category `gorm:"foreignKey:CategoryID"` // GORM association to the Category model.
	Blocked             bool      `gorm:"default:false"`         // A flag indicating whether the product is currently active (false) or inactive (true).
	Deleted             bool      `gorm:"default:false"`         // A flag indicating whether the product is currently active (false) or deleted (true).
}

// BeforeSave is a GORM hook that runs before a record is created or updated.
// It is used here to validate the product's price based on custom rules.
func (p *Product) BeforeSave(tx *gorm.DB) error {
	// Rule 1: The price must be greater than or equal to 100 cents.
	if p.PricePerUnitInCents < 100 {
		return errors.New("product price must be at least 100 cents")
	}

	// Rule 3: The price must be a multiple of 100.
	if p.PricePerUnitInCents%100 != 0 {
		return errors.New("product price must be in increments of 100 cents")
	}

	// If all validation rules pass, you can then proceed with other business logic checks.
	// For example, ensuring the category belongs to the same shop.

	return nil // Return nil if all validation passes, allowing the save operation to proceed.
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
	return int64(p.PricePerUnitInCents)
}

// Returns a boolean indicating whether the product is currently blocked. 🛑
func (p Product) GetBlocked() bool {
	return p.Blocked
}

// Returns a pointer to the unique ID of the category where the product belongs to
func (p Product) GetCategoryID() *int64 {
	if p.CategoryID == nil {
		return nil
	}
	id := int64(*p.CategoryID)
	return &id
}

// Returns a boolean indicating whether the product is currently deleted.
func (p Product) GetDeleted() bool {
	return p.Deleted
}

// Returns the unique identifier of the shop associated with this product. 🏪
func (p Product) GetShopID() int64 {
	return int64(p.ShopID)
}