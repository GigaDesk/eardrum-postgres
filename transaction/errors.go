package transaction

import (
    "fmt"
    "net/http"
    
    // Assuming this import path for your central error struct
    customErrors "github.com/GigaDesk/eardrum-interfaces/errors" 
)

// --- 4XX CLIENT/BUSINESS LOGIC ERRORS ---

// NewUnauthorizedError returns a 401 Unauthorized error.
// Used when credentials provided are invalid or missing (e.g., token check).
func NewUnauthorizedError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusUnauthorized, 
        message, 
        nil, 
    )
}

// NewPaymentRequiredError returns a 402 Payment Required error.
// This is the dedicated status code for insufficient funds or similar financial blocks.
func NewPaymentRequiredError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusPaymentRequired, // 402
        message, 
        nil, 
    )
}

// NewForbiddenFailure returns a 403 Forbidden error.
// Now reserved for other business logic blocks (e.g., product is deleted, merchant mismatch).
func NewForbiddenFailure(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusForbidden, // 403
        message, 
        nil, 
    )
}

// ErrTransactionFailed returns a 400 Bad Request error.
// Used for simple request failures (e.g., invalid amount, missing fields, zero or negative transfer amount).
func ErrTransactionFailed(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusBadRequest, 
        message, 
        nil, 
    )
}

// --- 5XX SERVER/SYSTEM FAILURES ---

// ErrDBLookupFailure returns a 500 Internal Server Error for failures during read queries.
// Wraps the underlying error for logging.
func ErrDBLookupFailure(message string, err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        message, 
        fmt.Errorf("db lookup error in transaction: %w", err),
    )
}

// ErrDBPersistenceFailure returns a 500 Internal Server Error for failures during write operations.
// Used for unexpected database errors (deadlocks, connection loss, commit failure, etc.).
func ErrDBPersistenceFailure(err error) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        "A critical database persistence operation failed within the transaction.",
        fmt.Errorf("db persistence error in transaction: %w", err),
    )
}

// NewAccountNotFoundError returns a 404 Not Found error.
// Used when a specific account identifier (UID, card number, etc.) does not exist in the system.
func NewAccountNotFoundError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusNotFound, // 404
        message,
        nil, 
    )
}