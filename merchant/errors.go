package merchant

import (
    "fmt"
    "net/http" // We must import net/http here to get the status constants
    
    // Import the central PublicError constructor
    customErrors "github.com/GigaDesk/eardrum-interfaces/errors" 
)

// ErrMerchantNotFound returns a 404 Not Found error.
func ErrMerchantNotFound(lookupKey string, lookupValue interface{}) *customErrors.PublicError {
    message := fmt.Sprintf("Merchant not found with %s: %v.", lookupKey, lookupValue)
    // Relies on the basic NewHTTPError(404, message, nil)
    return customErrors.NewHTTPError(
        http.StatusNotFound, 
        message, 
        nil, 
    )
}

// ErrMerchantConflict returns a 409 Conflict error.
func ErrMerchantConflict(message string, err error) *customErrors.PublicError {
    // Relies on the basic NewHTTPError(409, message, err)
    return customErrors.NewHTTPError(
        http.StatusConflict, 
        fmt.Sprintf("Merchant creation failed: %s", message),
        err, 
    )
}

// ErrDBLookupFailure returns a 500 Internal Server Error for failures during SELECT queries.
func ErrDBLookupFailure(message string, err error) *customErrors.PublicError {
    // Relies on the basic NewHTTPError(500, message, err)
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        message, 
        fmt.Errorf("db lookup error: %w", err),
    )
}

// ErrDBPersistenceFailure returns a 500 Internal Server Error for failures during INSERT/UPDATE/DELETE.
func ErrDBPersistenceFailure(err error) *customErrors.PublicError {
    // Relies on the basic NewHTTPError(500, message, err)
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        "A critical database operation (insert/update) failed unexpectedly for merchant data.",
        fmt.Errorf("db persistence error: %w", err),
    )
}

// NOTE: This must check the actual database driver error code (e.g., Postgres "23505").
// For now, it is a placeholder.
func isUniqueConstraintViolation(err error) bool {
    // In a real application, you would inspect the underlying driver error (e.g., *pq.Error or *mysql.MySQLError)
    // to check for the specific unique constraint code.
    return false 
}