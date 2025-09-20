package transaction

import (
	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"gorm.io/gorm"
)

// Gets a purchase by its unique id
func GetPurchaseWithId(Db *gorm.DB, Id int) (purchase.Purchase, error) {
	var purchase *Purchase
	//fetch the record from the database
	if err := Db.First(&purchase, Id).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}

// Gets all the purchases registered in the database
func GetPurchases(Db *gorm.DB) ([]purchase.Purchase, error) {

	var purchases []*Purchase

	if err := Db.Find(&purchases).Error; err != nil {
		return nil, err
	}

	var purchaselist []purchase.Purchase

	for _, p := range purchases {
		purchaselist = append(purchaselist, p)
	}

	return purchaselist, nil
}

// GetPurchasesForTransaction retrieves all purchases for a specific transaction.
// It uses the foreign key to filter the purchase records.
func GetPurchasesForTransaction(db *gorm.DB, transactionID uint) ([]purchase.Purchase, error) {
	var purchases []Purchase
	if err := db.Where("transaction_id = ?", transactionID).Find(&purchases).Error; err != nil {
		return nil, err
	}

	var purchaselist []purchase.Purchase

	for _, p := range purchases {
		purchaselist = append(purchaselist, p)
	}

	return purchaselist, nil
}

// GetPurchasesForProduct retrieves all purchases for a specific product.
// It uses the foreign key to filter the purchase records.
func GetPurchasesForProduct(db *gorm.DB, productID uint) ([]purchase.Purchase, error) {
	var purchases []Purchase
	if err := db.Where("product_id = ?", productID).Find(&purchases).Error; err != nil {
		return nil, err
	}
	var purchaselist []purchase.Purchase

	for _, p := range purchases {
		purchaselist = append(purchaselist, p)
	}

	return purchaselist, nil
}
