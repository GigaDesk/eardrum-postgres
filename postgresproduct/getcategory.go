package postgresproduct

import (
	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

//Gets a category by its unique id
func GetCategoryWithId(Db *gorm.DB, Id int) (product.Category, error) {
	var category *Category
	//fetch the record from the database
	if err := Db.First(&category, Id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

//Gets all the categories registered in the database
func GetCategories(Db *gorm.DB) ([]product.Category, error) {

	var categories []*Category

	if err := Db.Find(&categories).Error; err != nil {
		return nil, err
	}

	var categorylist []product.Category

	for _, c := range categories {
		categorylist = append(categorylist, c)
	}

	return categorylist, nil
}

// GetCategoriesForShop retrieves all categories for a given shop ID.
// It returns a slice of Category and an error if the query fails.
func GetCategoriesForShop(db *gorm.DB, shopID uint) ([]product.Category, error) {
    var categories []Category
    // Use the `Where` method to filter categories by their `ShopID`.
    // The `Find` method will populate the `categories` slice with the results.
    if err := db.Where("shop_id = ?", shopID).Find(&categories).Error; err != nil {
        return nil, err
    }

	var categorylist []product.Category

	for _, c := range categories {
		categorylist = append(categorylist, c)
	}

    return categorylist, nil
}

// GetBlockedCategoriesByShop fetches all categories that are marked as blocked for a specific shop.
func GetBlockedCategoriesByShop(db *gorm.DB, shopID uint) ([]Category, error) {
	var categories []Category
	// Find all categories where the shop ID matches and the 'blocked' flag is true.
	if err := db.Where("shop_id = ? AND blocked = ?", shopID, true).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}