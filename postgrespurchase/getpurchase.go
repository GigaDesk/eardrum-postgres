package postgrespurchase

import (
	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"gorm.io/gorm"
)

//Gets a purchase by its unique id
func GetPurchaseWithId(Db *gorm.DB, Id int) (purchase.Purchase, error) {
	var purchase *Purchase
	//fetch the record from the database
	if err := Db.First(&purchase, Id).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}

//Gets all the purchases registered in the database
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