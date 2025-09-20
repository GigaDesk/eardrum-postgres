package product

import (
	"errors"
	"gorm.io/gorm"
)

// DeleteCategory performs a hard delete on a category and sets the CategoryID to NULL for all
// associated products in a single database transaction to ensure data integrity.
// It now takes a shopID to validate ownership.
func DeleteCategory(db *gorm.DB, shopID uint, categoryID uint) error {
    // Start a new database transaction. All operations within the function passed to
    // db.Transaction will either be committed together or rolled back if an error occurs.
    err := db.Transaction(func(tx *gorm.DB) error {
        // First, update all products that belong to the specified category AND shop.
        // We set their CategoryID to NULL (represented by `nil` in Go) to unlink them.
        if err := tx.Model(&Product{}).Where("category_id = ? AND shop_id = ?", categoryID, shopID).Update("category_id", nil).Error; err != nil {
            // If the update fails, return the error. This will cause the transaction to roll back.
            return err
        }

        // Second, permanently delete the category record, ensuring it belongs to the correct shop.
        // The `Unscoped()` method bypasses GORM's soft-delete filter, performing a hard delete.
        if err := tx.Unscoped().Where("id = ? AND shop_id = ?", categoryID, shopID).Delete(&Category{}).Error; err != nil {
            // If the delete fails, return the error. This will cause the transaction to roll back.
            return err
        }

        // If both operations are successful, return nil to commit the transaction.
        return nil
    })
    
    // Return any error from the transaction.
    return err
}

// BlockCategoryWithProducts sets the 'blocked' flag to true for a category and all
// products associated with it in a single, atomic transaction.
func BlockCategoryWithProducts(db *gorm.DB, shopID, categoryID uint) error {
	// Start a transaction to ensure both updates happen together or fail together.
	err := db.Transaction(func(tx *gorm.DB) error {
		// First, block the category itself.
		// Use `Model` on `Category` and `Where` to apply the update only to the
		// specified category ID and shop ID.
		result := tx.Model(&Category{}).
			Where("id = ? AND shop_id = ?", categoryID, shopID).
			Update("blocked", true)

		if result.Error != nil {
			return result.Error
		}
		// Check if any rows were affected to ensure the category existed and belonged to the shop.
		if result.RowsAffected == 0 {
			return errors.New("category not found or does not belong to the specified shop")
		}

		// Second, block all products in that category.
		// Use `Model` on `Product` and `Where` to apply the update to all products
		// with the specified category ID and shop ID.
		result = tx.Model(&Product{}).
			Where("category_id = ? AND shop_id = ?", categoryID, shopID).
			Update("blocked", true)

		if result.Error != nil {
			return result.Error
		}

		// Return nil if both operations are successful, committing the transaction.
		return nil
	})

	return err
}

// AddProductsToCategory updates an array of products to a specific category.
// It sets the CategoryID for each product in the provided slice of productIDs.
// The operation is performed within a transaction to ensure data integrity.
func AddProductsToCategory(db *gorm.DB, shopID uint, categoryID uint, productIDs []uint) error {
    // Start a new database transaction.
    err := db.Transaction(func(tx *gorm.DB) error {
        // First, check if the category exists AND belongs to the specified shop.
        var category Category
        if err := tx.Where("id = ? AND shop_id = ?", categoryID, shopID).First(&category).Error; err != nil {
            // If the category is not found or does not belong to the shop, return the error.
            return err
        }

        // Next, check if all the products exist AND belong to the specified shop.
        var products []Product
        if err := tx.Where("id IN ? AND shop_id = ?", productIDs, shopID).Find(&products).Error; err != nil {
            return err
        }

        // Check if the number of products found matches the number of product IDs provided.
        // This validates that all products exist and belong to the correct shop.
        if len(products) != len(productIDs) {
            return errors.New("one or more products not found or do not belong to the specified shop")
        }

        // Use `Model` to specify the `Product` table and `Where` with `IN` to update all
        // products whose IDs are in the provided slice.
        // We set the `CategoryID` to the provided `categoryID`.
        if err := tx.Model(&Product{}).Where("id IN ?", productIDs).Update("category_id", categoryID).Error; err != nil {
            // If the update fails, return the error to trigger a rollback.
            return err
        }

        // If all checks pass and the product update is successful, return nil to commit.
        return nil
    })

    // Return any error from the transaction.
    return err
}


func RemoveProductsFromCategory(db *gorm.DB, shopID uint, productIDs []uint) error {
	// Update the 'CategoryID' field to nil (NULL in the database) for the specified products.
	// We use the `shopID` and `productIDs` to ensure security and target the correct products.
	result := db.Model(&Product{}).
		Where("id IN ? AND shop_id = ?", productIDs, shopID).
		Update("category_id", nil)

	if result.Error != nil {
		return result.Error
	}

	// It's a good practice to check if any rows were affected to verify the products existed
	// and belonged to the shop.
	if result.RowsAffected == 0 {
		return errors.New("no products found to remove from category or they do not belong to the specified shop")
	}

	return nil
}

// UpdateCategory updates one or more fields of a category, ensuring it belongs to the specified shop.
// This is a single, secure, and flexible function that replaces the two redundant functions.
func UpdateCategory(db *gorm.DB, shopID uint, categoryID uint, updates map[string]interface{}) (Category, error) {
    var category Category
    // Start a transaction to ensure the find and update are atomic.
    err := db.Transaction(func(tx *gorm.DB) error {
        // Find the category by ID and shopID to ensure it exists and belongs to the shop.
        if err := tx.Where("id = ? AND shop_id = ?", categoryID, shopID).First(&category).Error; err != nil {
            // If not found, return an error. The transaction will be rolled back.
            return err
        }

        // Update the category with the provided map of fields.
        if err := tx.Model(&category).Updates(updates).Error; err != nil {
            return err
        }
        return nil
    })

    // Return the updated category and any error from the transaction.
    return category, err
}
