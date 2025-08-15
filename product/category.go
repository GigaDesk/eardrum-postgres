package product

import (
	"time"

	"github.com/GigaDesk/eardrum-interfaces/shop"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"size:100;not null;uniqueIndex:idx_category_name_shop"` // Name of the category, e.g., "Electronics" or "Groceries." This is unique per shop.
	Description string `gorm:"type:text;not null"`                                // A detailed description of the category's purpose or contents.
	Blocked     bool      `gorm:"default:false"`                                        // A flag indicating whether the category is currently active (false) or inactive (true).
	ShopID      uint      `gorm:"not null;uniqueIndex:idx_category_name_shop"`          // The foreign key linking this category to its parent shop.
	Shop        shop.Shop `gorm:"foreignKey:ShopID"`                                    // GORM association to the Shop model, enabling database joins.
}

// Returns the unique identifier of the product category. 🆔
func (c Category) GetID() int64 {
	return int64(c.ID)
}

// Returns the creation timestamp of the product category in Coordinated Universal Time (UTC). 🗓️
func (c Category) GetCreatedAt() time.Time {
	return c.CreatedAt.UTC()
}

// Returns the last update timestamp of the product category in Coordinated Universal Time (UTC). ✍️
func (c Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt.UTC()
}

// Returns the timestamp of when the product category was soft-deleted, in Coordinated Universal Time (UTC). 🗑️
func (c Category) GetDeletedAt() time.Time {
	return c.DeletedAt.Time.UTC()
}

// Returns the name of the product category, e.g., "Electronics" or "Groceries." 🏷️
func (c Category) GetName() string {
	return c.Name
}

// Returns the descriptive text for the product category. 📝
func (c Category) GetDescription() string {
	return c.Description
}

// Returns a boolean indicating whether the product category is currently blocked. 🛑
func (c Category) GetBlocked() bool {
	return c.Blocked
}

// Returns the unique identifier of the shop associated with this category. 🏪
func (c Category) GetShopID() int64 {
	return int64(c.ShopID)
}
