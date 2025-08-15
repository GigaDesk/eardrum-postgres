package postgresproduct

import (
	"errors"

	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

//Gets a product by its unique id
func GetProductWithId(Db *gorm.DB, Id int) (product.Product, error) {
	var product *Product
	//fetch the record to be updated from the database
	if err := Db.First(&product, Id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

//Gets all the products registered in the database
func GetProducts(Db *gorm.DB) ([]product.Product, error) {

	var products []*Product

	if err := Db.Find(&products).Error; err != nil {
		return nil, err
	}

	var productlist []product.Product

	for _, p := range products {
		productlist = append(productlist, p)
	}

	return productlist, nil
}


// GetProductsForCategory retrieves all products for a given category ID that have not been soft-deleted.
// It returns a slice of Product and an error if the query fails.
func GetProductsForCategory(db *gorm.DB, categoryID *uint) ([]product.Product, error) {
    // Check if the provided categoryID is nil. If so, return an error.
    if categoryID == nil {
        return nil, errors.New("category ID cannot be nil")
    }
    
    // Create an empty slice of the GORM `Product` model to hold the results.
    var products []Product
    
    // Filter products by category_id and where the custom "Deleted" field is false.
    if err := db.Where("category_id = ? AND deleted = ?", categoryID, false).Find(&products).Error; err != nil {
        return nil, err
    }
    
    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}

// GetProductsForShop retrieves all products for a given shop ID that have not been soft-deleted.
// It returns a slice of Product and an error if the query fails.
func GetProductsForShop(db *gorm.DB, shopID uint) ([]product.Product, error) {
	var products []Product
    // Filter products by shop_id and where the custom "Deleted" field is false.
    if err := db.Where("shop_id = ? AND deleted = ?", shopID, false).Find(&products).Error; err != nil {
		return nil, err
	}

	var productList []product.Product

	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetUncategorizedProducts retrieves all products that do not have a category assigned (CategoryID is NULL),
// and have not been deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedProducts(db *gorm.DB) ([]product.Product, error) {
    var products []Product
    // Explicitly check for a NULL category_id and where the custom "Deleted" field is false.
    if err := db.Where("category_id IS NULL AND deleted = ?", false).Find(&products).Error; err != nil {
        return nil, err
    }

    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}

// GetUncategorizedProductsForShop retrieves all uncategorized products for a specific shop,
// filtering out any that have been deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedProductsForShop(db *gorm.DB, shopID uint) ([]product.Product, error) {
    var products []Product
    // Use chained `Where` methods to filter by `ShopID`, a NULL `CategoryID`, and where "Deleted" is false.
    if err := db.Where("shop_id = ? AND category_id IS NULL AND deleted = ?", shopID, false).Find(&products).Error; err != nil {
        return nil, err
    }

    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}


// =====================================================================================================================
// NEW FUNCTIONS FOR BLOCKED PRODUCTS
// =====================================================================================================================

// GetActiveProductsForCategory retrieves all products for a given category ID that are not blocked or deleted.
// It returns a slice of Product and an error if the query fails.
func GetActiveProductsForCategory(db *gorm.DB, categoryID *uint) ([]product.Product, error) {
    if categoryID == nil {
        return nil, errors.New("category ID cannot be nil")
    }
    
    var products []Product
    // Filter for a specific category, and where both the "Deleted" and "Blocked" fields are false.
    if err := db.Where("category_id = ? AND deleted = ? AND blocked = ?", categoryID, false, false).Find(&products).Error; err != nil {
        return nil, err
    }
    
    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}

// GetActiveProductsForShop retrieves all products for a given shop ID that are not blocked or deleted.
// It returns a slice of Product and an error if the query fails.
func GetActiveProductsForShop(db *gorm.DB, shopID uint) ([]product.Product, error) {
	var products []Product
    // Filter for a specific shop, and where both the "Deleted" and "Blocked" fields are false.
    if err := db.Where("shop_id = ? AND deleted = ? AND blocked = ?", shopID, false, false).Find(&products).Error; err != nil {
		return nil, err
	}

	var productList []product.Product

	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetUncategorizedActiveProducts retrieves all products that are not categorized, blocked, or deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedActiveProducts(db *gorm.DB) ([]product.Product, error) {
    var products []Product
    // Filter for uncategorized products, and where both the "Deleted" and "Blocked" fields are false.
    if err := db.Where("category_id IS NULL AND deleted = ? AND blocked = ?", false, false).Find(&products).Error; err != nil {
        return nil, err
    }

    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}

// GetUncategorizedActiveProductsForShop retrieves all uncategorized products for a specific shop
// that are not blocked or deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedActiveProductsForShop(db *gorm.DB, shopID uint) ([]product.Product, error) {
    var products []Product
    // Filter for a specific shop, uncategorized products, and where both the "Deleted" and "Blocked" fields are false.
    if err := db.Where("shop_id = ? AND category_id IS NULL AND deleted = ? AND blocked = ?", shopID, false, false).Find(&products).Error; err != nil {
        return nil, err
    }

    var productList []product.Product
    for _, p := range products {
        productList = append(productList, p)
    }

    return productList, nil
}

// GetDeletedProductsByShop fetches all products that are marked as deleted for a specific shop.
func GetDeletedProductsByShop(db *gorm.DB, shopID uint) ([]Product, error) {
	var products []Product
	// Find all products where the shop ID matches and the 'deleted' flag is true.
	if err := db.Where("shop_id = ? AND deleted = ?", shopID, true).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}


// GetBlockedProductsByShop fetches all products that are marked as blocked for a specific shop.
func GetBlockedProductsByShop(db *gorm.DB, shopID uint) ([]Product, error) {
	var products []Product
	// Find all products where the shop ID matches and the 'blocked' flag is true.
	if err := db.Where("shop_id = ? AND blocked = ?", shopID, true).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}