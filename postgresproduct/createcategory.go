package postgresproduct

import (
	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

//create a category record
func CreateCategory(c product.NewCategory, Db *gorm.DB) (product.Category, error) {
	//create category data
	category := &Category{
		Name:        c.GetName(),
		Description: c.GetDescription(),
	}

	//create a category record in the database and return if operation succeeds
	if err := Db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}