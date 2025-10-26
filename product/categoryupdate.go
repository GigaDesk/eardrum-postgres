package product

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// DeleteCategory performs a hard delete on a category and sets the CategoryID to NULL for all
// associated products in a single database transaction to ensure data integrity.
// It now takes a merchantID to validate ownership.
func DeleteCategory(db *gorm.DB, merchantID uint, categoryID uint) error {
    // Start a new database transaction. All operations within the function passed to
    // db.Transaction will either be committed together or rolled back if an error occurs.
    err := db.Transaction(func(tx *gorm.DB) error {
        // First, update all products that belong to the specified category AND merchant.
        // We set their CategoryID to NULL (represented by `nil` in Go) to unlink them.
        if err := tx.Model(&Product{}).Where("category_id = ? AND merchant_id = ?", categoryID, merchantID).Update("category_id", nil).Error; err != nil {
            // If the update fails, return the error. This will cause the transaction to roll back.
            // Use 500 Persistence Failure for unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to unlink products from category %d: %w", categoryID, err))
        }

        // Second, permanently delete the category record, ensuring it belongs to the correct merchant.
        // The `Unscoped()` method bypasses GORM's soft-delete filter, performing a hard delete.
        result := tx.Unscoped().Where("id = ? AND merchant_id = ?", categoryID, merchantID).Delete(&Category{})
        if result.Error != nil {
            // If the delete fails, return the error. This will cause the transaction to roll back.
            // Use 500 Persistence Failure for unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to delete category %d: %w", categoryID, result.Error))
        }

        // Check if any rows were affected. If not, the category either didn't exist or didn't belong to the merchant.
        if result.RowsAffected == 0 {
            // Use 404 Not Found error.
            return ErrCategoryNotFound("id", categoryID)
        }

        // If both operations are successful, return nil to commit the transaction.
        return nil
    })
    
    // Return any error from the transaction.
    return err
}

// BlockCategoryWithProducts sets the 'blocked' flag to true for a category and all
// products associated with it in a single, atomic transaction.
func BlockCategoryWithProducts(db *gorm.DB, merchantID, categoryID uint) error {
    // Start a transaction to ensure both updates happen together or fail together.
    err := db.Transaction(func(tx *gorm.DB) error {
        // First, block the category itself.
        // Use `Model` on `Category` and `Where` to apply the update only to the
        // specified category ID and merchant ID.
        result := tx.Model(&Category{}).
            Where("id = ? AND merchant_id = ?", categoryID, merchantID).
            Update("blocked", true)

        if result.Error != nil {
            // Use 500 Persistence Failure for unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to block category %d: %w", categoryID, result.Error))
        }
        // Check if any rows were affected to ensure the category existed and belonged to the merchant.
        if result.RowsAffected == 0 {
            // Use 404 Not Found error.
            return ErrCategoryNotFound("id", categoryID)
        }

        // Second, block all products in that category.
        // Use `Model` on `Product` and `Where` to apply the update to all products
        // with the specified category ID and merchant ID.
        result = tx.Model(&Product{}).
            Where("category_id = ? AND merchant_id = ?", categoryID, merchantID).
            Update("blocked", true)

        if result.Error != nil {
            // Use 500 Persistence Failure for unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to block products in category %d: %w", categoryID, result.Error))
        }

        // Return nil if both operations are successful, committing the transaction.
        return nil
    })

    return err
}

// AddProductsToCategory updates an array of products to a specific category.
// It sets the CategoryID for each product in the provided slice of productIDs.
// The operation is performed within a transaction to ensure data integrity.
func AddProductsToCategory(db *gorm.DB, merchantID uint, categoryID uint, productIDs []uint) error {
    // Start a new database transaction.
    err := db.Transaction(func(tx *gorm.DB) error {
        // First, check if the category exists AND belongs to the specified merchant.
        var category Category
        categoryLookupErr := tx.Where("id = ? AND merchant_id = ?", categoryID, merchantID).First(&category).Error
        if categoryLookupErr != nil {
            // If the category is not found or does not belong to the merchant, return the error.
            if errors.Is(categoryLookupErr, gorm.ErrRecordNotFound) {
                // Use 404 Not Found error.
                return ErrCategoryNotFound("id", categoryID)
            }
            // Use 500 Lookup Failure for unexpected DB read errors.
            return ErrDBLookupFailure(fmt.Sprintf("Failed to lookup category %d.", categoryID), categoryLookupErr)
        }

        // Next, check if all the products exist AND belong to the specified merchant.
        var products []Product
        productLookupErr := tx.Where("id IN ? AND merchant_id = ?", productIDs, merchantID).Find(&products).Error
        if productLookupErr != nil {
            // Use 500 Lookup Failure for unexpected DB read errors.
            return ErrDBLookupFailure("Failed to lookup products for category addition.", productLookupErr)
        }

        // Check if the number of products found matches the number of product IDs provided.
        // This validates that all products exist and belong to the correct merchant.
        if len(products) != len(productIDs) {
            // One or more products were missing or didn't match the merchant. Use a 404/Not Found error.
            return ErrProductNotFound("one or more IDs", productIDs)
        }

        // Use `Model` to specify the `Product` table and `Where` with `IN` to update all
        // products whose IDs are in the provided slice.
        // We set the `CategoryID` to the provided `categoryID`.
        if err := tx.Model(&Product{}).Where("id IN ?", productIDs).Update("category_id", categoryID).Error; err != nil {
            // If the update fails, return the error to trigger a rollback.
            // Use 500 Persistence Failure for unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to add products to category %d: %w", categoryID, err))
        }

        // If all checks pass and the product update is successful, return nil to commit.
        return nil
    })

    // Return any error from the transaction.
    return err
}


func RemoveProductsFromCategory(db *gorm.DB, merchantID uint, productIDs []uint) error {
    // Update the 'CategoryID' field to nil (NULL in the database) for the specified products.
    // We use the `merchantID` and `productIDs` to ensure security and target the correct products.
    result := db.Model(&Product{}).
        Where("id IN ? AND merchant_id = ?", productIDs, merchantID).
        Update("category_id", nil)

    if result.Error != nil {
        // Use 500 Persistence Failure for unexpected DB write errors.
        return ErrDBPersistenceFailure(fmt.Errorf("failed to remove products from category: %w", result.Error))
    }

    // It's a good practice to check if any rows were affected to verify the products existed
    // and belonged to the merchant.
    if result.RowsAffected == 0 {
        // Use 404 Not Found error since none of the specified products could be found 
        // belonging to the merchant or in the provided list.
        return ErrProductNotFound("one or more IDs", productIDs)
    }

    return nil
}

// UpdateCategory updates one or more fields of a category, ensuring it belongs to the specified merchant.
// This is a single, secure, and flexible function that replaces the two redundant functions.
func UpdateCategory(db *gorm.DB, merchantID uint, categoryID uint, updates map[string]interface{}) (Category, error) {
    var category Category
    // Start a transaction to ensure the find and update are atomic.
    err := db.Transaction(func(tx *gorm.DB) error {
        // Find the category by ID and merchantID to ensure it exists and belongs to the merchant.
        lookupErr := tx.Where("id = ? AND merchant_id = ?", categoryID, merchantID).First(&category).Error
        if lookupErr != nil {
            // If not found, return an error. The transaction will be rolled back.
            if errors.Is(lookupErr, gorm.ErrRecordNotFound) {
                 // Use 404 Not Found error.
                return ErrCategoryNotFound("id", categoryID)
            }
            // Use 500 Lookup Failure for unexpected DB read errors.
            return ErrDBLookupFailure(fmt.Sprintf("Failed to lookup category %d for update.", categoryID), lookupErr)
        }

        // Update the category with the provided map of fields.
        if err := tx.Model(&category).Updates(updates).Error; err != nil {
            // Check for conflict (assuming `isUniqueConstraintViolation` is implemented correctly).
            if isUniqueConstraintViolation(err) {
                 // Use 409 Conflict error.
                return ErrCategoryConflict(fmt.Sprintf("Update conflicted on a unique field for category %d.", categoryID), err)
            }
            // Use 500 Persistence Failure for all other unexpected DB write errors.
            return ErrDBPersistenceFailure(fmt.Errorf("failed to update category %d: %w", categoryID, err))
        }
        return nil
    })

    // Return the updated category and any error from the transaction.
    return category, err
}
