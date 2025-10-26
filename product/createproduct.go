package product

import (
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

// create a product record
func CreateProduct(p product.NewProduct, Db *gorm.DB, merchant_id uint) (product.Product, error) {
    //create product data
    product := &Product{
        Name:                p.GetName(),
        PricePerUnitInCents: uint(p.GetPricePerUnitInCents()),
        MerchantID:              merchant_id,
        // Note: CategoryID is omitted for basic product creation.
    }

    //create a product record in the database and return if operation succeeds
    if err := Db.Create(product).Error; err != nil {
        // Check for specific unique constraint violation (409 Conflict).
        if isUniqueConstraintViolation(err) {
            // Use 409 Conflict error helper for domain-specific naming/SKU issue.
            return nil, ErrProductConflict(fmt.Sprintf("Product with name '%s' already exists for merchant %d.", product.Name, merchant_id), err)
        }
        
        // If it's any other persistence error, return 500 Internal Server Error.
        return nil, ErrDBPersistenceFailure(fmt.Errorf("failed to create product record: %w", err))
    }

    // The concrete *Product struct satisfies the product.Product interface.
    return product, nil
}
