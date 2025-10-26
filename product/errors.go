package product

import (
    "fmt"
    "net/http" // We must import net/http here to get the status constants
    
    // Import the central PublicError constructor
    customErrors "github.com/GigaDesk/eardrum-interfaces/errors" 
)

// --- 4XX CLIENT/BUSINESS LOGIC ERRORS ---

// ErrProductNotFound returns a 404 Not Found error specifically for Product records.
func ErrProductNotFound(lookupKey string, lookupValue interface{}) *customErrors.PublicError {
    message := fmt.Sprintf("Product not found with %s: %v.", lookupKey, lookupValue)
    return customErrors.NewHTTPError(
        http.StatusNotFound, 
        message, 
        nil, 
    )
}

// ErrCategoryNotFound returns a 404 Not Found error specifically for Category records.
func ErrCategoryNotFound(lookupKey string, lookupValue interface{}) *customErrors.PublicError {
    message := fmt.Sprintf("Category not found with %s: %v.", lookupKey, lookupValue)
    return customErrors.NewHTTPError(
        http.StatusNotFound, 
        message, 
        nil, 
    )
}

// ErrProductConflict returns a 409 Conflict error specifically for product data.
// Used when creating or updating a Product record that violates a unique constraint (e.g., SKU/Name).
func ErrProductConflict(message string, err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusConflict, 
        fmt.Sprintf("Product data conflict: %s", message),
        err, 
    )
}

// ErrCategoryConflict returns a 409 Conflict error specifically for category data.
// Used when creating or updating a Category record that violates a unique constraint (e.g., Name).
func ErrCategoryConflict(message string, err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusConflict, 
        fmt.Sprintf("Category data conflict: %s", message),
        err, 
    )
}

// --- 5XX SERVER/SYSTEM FAILURES (GENERIC) ---

// ErrDBLookupFailure returns a 500 Internal Server Error for failures during SELECT queries.
// This helper is generic and should be used for any unexpected database read failure.
func ErrDBLookupFailure(message string, err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        message, 
        fmt.Errorf("db lookup error: %w", err),
    )
}

// ErrDBPersistenceFailure returns a 500 Internal Server Error for failures during INSERT/UPDATE/DELETE.
// This helper is generic and should be used for any unexpected database write failure.
func ErrDBPersistenceFailure(err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        "A critical database persistence operation for product/category failed unexpectedly.",
        fmt.Errorf("db persistence error: %w", err),
    )
}

// NOTE: This utility function is required by CRUD logic to differentiate between a general 
// persistence failure (500) and a conflict/unique constraint violation (409).
func isUniqueConstraintViolation(err error) bool {
    // In a real application, you would inspect the underlying driver error (e.g., *pq.Error or *mysql.MySQLError)
    // to check for the specific unique constraint code (e.g., "23505" in Postgres).
    return false 
}
