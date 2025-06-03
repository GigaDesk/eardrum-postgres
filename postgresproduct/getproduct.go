package postgresproduct

import (
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