package product

import (
	"errors"
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

// Gets a product by its unique id
func GetProductWithId(Db *gorm.DB, Id int) (product.Product, error) {
	var p *Product
	//fetch the record to be updated from the database
	if err := Db.First(&p, Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Specific 404 error
			return nil, ErrProductNotFound("id", Id)
		}
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to look up product with ID %d.", Id), err)
	}

	return p, nil
}

// Gets all the products registered in the database
func GetProducts(Db *gorm.DB) ([]product.Product, error) {

	var products []*Product

	if err := Db.Find(&products).Error; err != nil {
		// General 500 system error for unexpected DB issues in bulk fetch
		return nil, ErrDBLookupFailure("Failed to retrieve all products.", err)
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
		// This is a logic error, use 500 for the data layer
		return nil, ErrDBLookupFailure("Category ID for lookup cannot be nil.", errors.New("nil category ID"))
	}

	// Create an empty slice of the GORM `Product` model to hold the results.
	var products []Product

	// Filter products by category_id and where the custom "Deleted" field is false.
	if err := db.Where("category_id = ? AND deleted = ?", *categoryID, false).Find(&products).Error; err != nil {
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve products for category ID %d.", *categoryID), err)
	}

	var productList []product.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetProductsForMerchant retrieves all products for a given merchant ID that have not been soft-deleted.
// It returns a slice of Product and an error if the query fails.
func GetProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter products by merchant_id and where the custom "Deleted" field is false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND deleted = ?", merchantID, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve products for merchant ID %d.", merchantID), err)
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
		// General 500 system error
		return nil, ErrDBLookupFailure("Failed to retrieve uncategorized products.", err)
	}

	var productList []product.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetUncategorizedProductsForMerchant retrieves all uncategorized products for a specific merchant, // CHANGED Shop to Merchant
// filtering out any that have been deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Use chained `Where` methods to filter by `MerchantID`, a NULL `CategoryID`, and where "Deleted" is false. // CHANGED ShopID to MerchantID
	if err := db.Where("merchant_id = ? AND category_id IS NULL AND deleted = ?", merchantID, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve uncategorized products for merchant ID %d.", merchantID), err)
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
		// Logic error check, use 500
		return nil, ErrDBLookupFailure("Category ID for active product lookup cannot be nil.", errors.New("nil category ID"))
	}

	var products []Product
	// Filter for a specific category, and where both the "Deleted" and "Blocked" fields are false.
	if err := db.Where("category_id = ? AND deleted = ? AND blocked = ?", *categoryID, false, false).Find(&products).Error; err != nil {
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve active products for category ID %d.", *categoryID), err)
	}

	var productList []product.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetActiveProductsForMerchant retrieves all products for a given merchant ID that are not blocked or deleted. // CHANGED Shop to Merchant
// It returns a slice of Product and an error if the query fails.
func GetActiveProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter for a specific merchant, and where both the "Deleted" and "Blocked" fields are false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND deleted = ? AND blocked = ?", merchantID, false, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve active products for merchant ID %d.", merchantID), err)
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
		// General 500 system error
		return nil, ErrDBLookupFailure("Failed to retrieve uncategorized active products.", err)
	}

	var productList []product.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetUncategorizedActiveProductsForMerchant retrieves all uncategorized products for a specific merchant // CHANGED Shop to Merchant
// that are not blocked or deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedActiveProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter for a specific merchant, uncategorized products, and where both the "Deleted" and "Blocked" fields are false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND category_id IS NULL AND deleted = ? AND blocked = ?", merchantID, false, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve uncategorized active products for merchant ID %d.", merchantID), err)
	}

	var productList []product.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	return productList, nil
}

// GetDeletedProductsByMerchant fetches all products that are marked as deleted for a specific merchant. // CHANGED Shop to Merchant
func GetDeletedProductsByMerchant(db *gorm.DB, merchantID uint) ([]Product, error) { // CHANGED Shop to Merchant and shopID to merchantID
	var products []Product
	// Find all products where the merchant ID matches and the 'deleted' flag is true. // CHANGED shop ID to merchant ID
	if err := db.Where("merchant_id = ? AND deleted = ?", merchantID, true).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve deleted products for merchant ID %d.", merchantID), err)
	}
	return products, nil
}


// GetBlockedProductsByMerchant fetches all products that are marked as blocked for a specific merchant. // CHANGED Shop to Merchant
func GetBlockedProductsByMerchant(db *gorm.DB, merchantID uint) ([]Product, error) { // CHANGED Shop to Merchant and shopID to merchantID
	var products []Product
	// Find all products where the merchant ID matches and the 'blocked' flag is true. // CHANGED shop ID to merchant ID
	if err := db.Where("merchant_id = ? AND blocked = ?", merchantID, true).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		// General 500 system error
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve blocked products for merchant ID %d.", merchantID), err)
	}
	return products, nil
}
