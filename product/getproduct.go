package product

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

// GetProductsForMerchant retrieves all products for a given merchant ID that have not been soft-deleted.
// It returns a slice of Product and an error if the query fails.
func GetProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter products by merchant_id and where the custom "Deleted" field is false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND deleted = ?", merchantID, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
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

// GetUncategorizedProductsForMerchant retrieves all uncategorized products for a specific merchant, // CHANGED Shop to Merchant
// filtering out any that have been deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Use chained `Where` methods to filter by `MerchantID`, a NULL `CategoryID`, and where "Deleted" is false. // CHANGED ShopID to MerchantID
	if err := db.Where("merchant_id = ? AND category_id IS NULL AND deleted = ?", merchantID, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
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

// GetActiveProductsForMerchant retrieves all products for a given merchant ID that are not blocked or deleted. // CHANGED Shop to Merchant
// It returns a slice of Product and an error if the query fails.
func GetActiveProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter for a specific merchant, and where both the "Deleted" and "Blocked" fields are false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND deleted = ? AND blocked = ?", merchantID, false, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
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

// GetUncategorizedActiveProductsForMerchant retrieves all uncategorized products for a specific merchant // CHANGED Shop to Merchant
// that are not blocked or deleted.
// It returns a slice of Product and an error if the query fails.
func GetUncategorizedActiveProductsForMerchant(db *gorm.DB, merchantID uint) ([]product.Product, error) { // CHANGED shopID to merchantID
	var products []Product
	// Filter for a specific merchant, uncategorized products, and where both the "Deleted" and "Blocked" fields are false. // CHANGED shop_id to merchant_id
	if err := db.Where("merchant_id = ? AND category_id IS NULL AND deleted = ? AND blocked = ?", merchantID, false, false).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		return nil, err
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
		return nil, err
	}
	return products, nil
}


// GetBlockedProductsByMerchant fetches all products that are marked as blocked for a specific merchant. // CHANGED Shop to Merchant
func GetBlockedProductsByMerchant(db *gorm.DB, merchantID uint) ([]Product, error) { // CHANGED Shop to Merchant and shopID to merchantID
	var products []Product
	// Find all products where the merchant ID matches and the 'blocked' flag is true. // CHANGED shop ID to merchant ID
	if err := db.Where("merchant_id = ? AND blocked = ?", merchantID, true).Find(&products).Error; err != nil { // CHANGED shop_id to merchant_id
		return nil, err
	}
	return products, nil
}