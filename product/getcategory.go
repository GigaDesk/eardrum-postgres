package product

import (
	"errors"
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

// Gets a category by its unique id
func GetCategoryWithId(Db *gorm.DB, Id int) (product.Category, error) {
	var category *Category
	//fetch the record from the database
	if err := Db.First(&category, Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Specific 404 error for the client
			return nil, ErrCategoryNotFound("id", Id)
		}
		// General 500 system error for unexpected DB issues
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to look up category with ID %d.", Id), err)
	}

	return category, nil
}

// Gets all the categories registered in the database
func GetCategories(Db *gorm.DB) ([]product.Category, error) {

	var categories []*Category

	if err := Db.Find(&categories).Error; err != nil {
		// General 500 system error for unexpected DB issues in bulk fetch
		return nil, ErrDBLookupFailure("Failed to retrieve all categories.", err)
	}

	var categorylist []product.Category

	for _, c := range categories {
		categorylist = append(categorylist, c)
	}

	return categorylist, nil
}

// GetCategoriesForMerchant retrieves all categories for a given merchant ID.
// It returns a slice of Category and an error if the query fails.
func GetCategoriesForMerchant(db *gorm.DB, merchantID uint) ([]product.Category, error) {
	var categories []Category
	// Use the `Where` method to filter categories by their `MerchantID`.
	// The `Find` method will populate the `categories` slice with the results.
	if err := db.Where("merchant_id = ?", merchantID).Find(&categories).Error; err != nil {
		// General 500 system error for unexpected DB issues in bulk fetch
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve categories for merchant ID %d.", merchantID), err)
	}

	var categorylist []product.Category

	for _, c := range categories {
		categorylist = append(categorylist, c)
	}

	return categorylist, nil
}

// GetBlockedCategoriesByMerchant fetches all categories that are marked as blocked for a specific merchant.
func GetBlockedCategoriesByMerchant(db *gorm.DB, merchantID uint) ([]Category, error) {
	var categories []Category
	// Find all categories where the merchant ID matches and the 'blocked' flag is true.
	if err := db.Where("merchant_id = ? AND blocked = ?", merchantID, true).Find(&categories).Error; err != nil {
		// General 500 system error for unexpected DB issues in bulk fetch
		return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve blocked categories for merchant ID %d.", merchantID), err)
	}
	return categories, nil
}