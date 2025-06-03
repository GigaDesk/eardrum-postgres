package postgresproduct

import (
	"github.com/GigaDesk/eardrum-interfaces/product"
	"gorm.io/gorm"
)

func UpdateCategoryName(Db *gorm.DB, newName string, id int) (product.Category, error) {
	var category *Category
	//fetch the record to be updated from the database
	if err := Db.First(&category, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&category).Updates(map[string]interface{}{"name": newName}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&category, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return category, nil
}

func UpdateCategoryDescription(Db *gorm.DB, newDescription string, id int) (product.Category, error) {
	var category *Category
	//fetch the record to be updated from the database
	if err := Db.First(&category, id).Error; err != nil {
		return nil, err
	}

	// Update the records' attributes with `map`
	if err := Db.Model(&category).Updates(map[string]interface{}{"description": newDescription}).Error; err != nil {
		return nil, err
	}

	//fetch the record again from the database, this time the updated version
	if err := Db.First(&category, id).Error; err != nil {
		return nil, err
	}

	//return the updated record
	return category, nil
}
