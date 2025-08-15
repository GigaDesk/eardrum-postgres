package product

import (
	"errors"

	"gorm.io/gorm"
)

// =====================================================================================================================
// NEW FUNCTIONS FOR UPDATING PRODUCT STATUS
// =====================================================================================================================

// DeleteProduct sets a product's 'deleted' flag to true, securely checking for shop ownership.
func DeleteProduct(db *gorm.DB, shopID, productID uint) error {
    // Update the 'Deleted' field, ensuring the product belongs to the correct shop.
    result := db.Model(&Product{}).Where("id = ? AND shop_id = ?", productID, shopID).Update("deleted", true)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("product not found or does not belong to the specified shop")
    }
    return nil
}

// BlockProduct sets a product's 'blocked' flag to true, securely checking for shop ownership.
func BlockProduct(db *gorm.DB, shopID, productID uint) error {
    // Update the 'Blocked' field, ensuring the product belongs to the correct shop.
    result := db.Model(&Product{}).Where("id = ? AND shop_id = ?", productID, shopID).Update("blocked", true)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("product not found or does not belong to the specified shop")
    }
    return nil
}

// RestoreProduct sets a product's 'deleted' flag back to false, securely checking for shop ownership.
func RestoreProduct(db *gorm.DB, shopID, productID uint) error {
    // Update the 'Deleted' field, ensuring the product belongs to the correct shop.
    result := db.Model(&Product{}).Where("id = ? AND shop_id = ?", productID, shopID).Update("deleted", false)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("product not found or does not belong to the specified shop")
    }
    return nil
}

// UnblockProduct sets a product's 'blocked' flag back to false, securely checking for shop ownership.
func UnblockProduct(db *gorm.DB, shopID, productID uint) error {
    // Update the 'Blocked' field, ensuring the product belongs to the correct shop.
    result := db.Model(&Product{}).Where("id = ? AND shop_id = ?", productID, shopID).Update("blocked", false)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("product not found or does not belong to the specified shop")
    }
    return nil
}

// UpdateProductPrice updates a product's price, securely checking for shop ownership.
func UpdateProductPrice(db *gorm.DB, shopID, productID uint, newPrice uint) error {
    // Update the 'PricePerUnitInCents' field, ensuring the product belongs to the correct shop.
    result := db.Model(&Product{}).Where("id = ? AND shop_id = ?", productID, shopID).Update("price_per_unit_in_cents", newPrice)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("product not found or does not belong to the specified shop")
    }
    return nil
}
