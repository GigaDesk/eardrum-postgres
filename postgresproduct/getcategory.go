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