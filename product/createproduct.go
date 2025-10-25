package product

import (
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
	}

	//create a product record in the database and return if operation succeeds
	if err := Db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
