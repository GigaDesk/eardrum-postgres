package product

import (
	"fmt"

	"gorm.io/gorm"
)

// =====================================================================================================================
// NEW FUNCTIONS FOR UPDATING PRODUCT STATUS
// =====================================================================================================================

// DeleteProduct sets a product's 'deleted' flag to true, securely checking for merchant ownership.
func DeleteProduct(db *gorm.DB, merchantID, productID uint) error { // CHANGED shopID to merchantID
	// Update the 'Deleted' field, ensuring the product belongs to the correct merchant.
	result := db.Model(&Product{}).Where("id = ? AND merchant_id = ?", productID, merchantID).Update("deleted", true) // CHANGED shop_id to merchant_id
	if result.Error != nil {
		// System error during persistence (500)
		return ErrDBPersistenceFailure(fmt.Errorf("failed to delete product %d: %w", productID, result.Error))
	}
	if result.RowsAffected == 0 {
		// Product not found or does not belong to merchant (404)
		return ErrProductNotFound("ID", productID)
	}
	return nil
}

// BlockProduct sets a product's 'blocked' flag to true, securely checking for merchant ownership.
func BlockProduct(db *gorm.DB, merchantID, productID uint) error { // CHANGED shopID to merchantID
	// Update the 'Blocked' field, ensuring the product belongs to the correct merchant.
	result := db.Model(&Product{}).Where("id = ? AND merchant_id = ?", productID, merchantID).Update("blocked", true) // CHANGED shop_id to merchant_id
	if result.Error != nil {
		// System error during persistence (500)
		return ErrDBPersistenceFailure(fmt.Errorf("failed to block product %d: %w", productID, result.Error))
	}
	if result.RowsAffected == 0 {
		// Product not found or does not belong to merchant (404)
		return ErrProductNotFound("ID", productID)
	}
	return nil
}

// RestoreProduct sets a product's 'deleted' flag back to false, securely checking for merchant ownership.
func RestoreProduct(db *gorm.DB, merchantID, productID uint) error { // CHANGED shopID to merchantID
	// Update the 'Deleted' field, ensuring the product belongs to the correct merchant.
	result := db.Model(&Product{}).Where("id = ? AND merchant_id = ?", productID, merchantID).Update("deleted", false) // CHANGED shop_id to merchant_id
	if result.Error != nil {
		// System error during persistence (500)
		return ErrDBPersistenceFailure(fmt.Errorf("failed to restore product %d: %w", productID, result.Error))
	}
	if result.RowsAffected == 0 {
		// Product not found or does not belong to merchant (404)
		return ErrProductNotFound("ID", productID)
	}
	return nil
}

// UnblockProduct sets a product's 'blocked' flag back to false, securely checking for merchant ownership.
func UnblockProduct(db *gorm.DB, merchantID, productID uint) error { // CHANGED shopID to merchantID
	// Update the 'Blocked' field, ensuring the product belongs to the correct merchant.
	result := db.Model(&Product{}).Where("id = ? AND merchant_id = ?", productID, merchantID).Update("blocked", false) // CHANGED shop_id to merchant_id
	if result.Error != nil {
		// System error during persistence (500)
		return ErrDBPersistenceFailure(fmt.Errorf("failed to unblock product %d: %w", productID, result.Error))
	}
	if result.RowsAffected == 0 {
		// Product not found or does not belong to merchant (404)
		return ErrProductNotFound("ID", productID)
	}
	return nil
}

// UpdateProductPrice updates a product's price, securely checking for merchant ownership.
func UpdateProductPrice(db *gorm.DB, merchantID, productID uint, newPrice uint) error { // CHANGED shopID to merchantID
	// Update the 'PricePerUnitInCents' field, ensuring the product belongs to the correct merchant.
	result := db.Model(&Product{}).Where("id = ? AND merchant_id = ?", productID, merchantID).Update("price_per_unit_in_cents", newPrice) // CHANGED shop_id to merchant_id
	if result.Error != nil {
		// System error during persistence (500)
		return ErrDBPersistenceFailure(fmt.Errorf("failed to update price for product %d: %w", productID, result.Error))
	}
	if result.RowsAffected == 0 {
		// Product not found or does not belong to merchant (404)
		return ErrProductNotFound("ID", productID)
	}
	return nil
}
