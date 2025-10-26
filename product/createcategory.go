package product

import (
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

// create a category record
func CreateCategory(c product.NewCategory, Db *gorm.DB, merchant_id uint) (product.Category, error) {
    //create category data
    category := &Category{
        Name:        c.GetName(),
        Description: c.GetDescription(),
        MerchantID:      merchant_id,
    }

    //create a category record in the database and return if operation succeeds
    if err := Db.Create(category).Error; err != nil {
        // Check for specific unique constraint violation (409 Conflict).
        if isUniqueConstraintViolation(err) {
            // Use 409 Conflict error helper for domain-specific naming issue.
            return nil, ErrCategoryConflict(fmt.Sprintf("Category with name '%s' already exists for merchant %d.", category.Name, merchant_id), err)
        }
        
        // If it's any other persistence error, return 500 Internal Server Error.
        return nil, ErrDBPersistenceFailure(fmt.Errorf("failed to create category record: %w", err))
    }

    // The concrete *Category struct satisfies the product.Category interface.
    return category, nil
}

