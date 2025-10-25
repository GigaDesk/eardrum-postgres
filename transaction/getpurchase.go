package transaction

import (
	"errors"
	"fmt"

	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"gorm.io/gorm"
)

// Gets a purchase by its unique id
func GetPurchaseWithId(Db *gorm.DB, Id int) (purchase.Purchase, error) {
    var p *Purchase
    //fetch the record from the database
    if err := Db.First(&p, Id).Error; err != nil {
        // Allow gorm.ErrRecordNotFound to be returned for 404 mapping at the service layer.
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve purchase with ID %d.", Id), err)
    }

    return p, nil
}

// Gets all the purchases registered in the database
func GetPurchases(Db *gorm.DB) ([]purchase.Purchase, error) {

    var purchases []*Purchase

    if err := Db.Order("created_at desc").Find(&purchases).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure("Failed to retrieve all purchases.", err)
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
    if err := db.Where("transaction_id = ?", transactionID).Order("created_at desc").Find(&purchases).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve purchases for transaction ID %d.", transactionID), err)
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
    if err := db.Where("product_id = ?", productID).Order("created_at desc").Find(&purchases).Error; err != nil {
        // General DB lookup failure -> 500 Internal Server Error
        return nil, ErrDBLookupFailure(fmt.Sprintf("Failed to retrieve purchases for product ID %d.", productID), err)
    }
    var purchaselist []purchase.Purchase

    for _, p := range purchases {
        purchaselist = append(purchaselist, p)
    }

    return purchaselist, nil
}
