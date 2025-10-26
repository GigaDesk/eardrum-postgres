package user

import (
	"net/http"
	"fmt"
	
	// Assuming this is your central error definition package
	customErrors "github.com/GigaDesk/eardrum-interfaces/errors" 
)

// ErrUserConflict returns a 409 Conflict error for existing users.
func ErrUserConflict(message string) *customErrors.PublicError {
	// Business logic errors (like validation/existence) often don't wrap a raw error,
	// but we still return the PublicError structure.
	return customErrors.NewHTTPError(
		http.StatusConflict, 
		message,
		nil, 
	)
}

// ErrDBPersistenceFailure is for unexpected errors during the actual Create or Update operation (500).
func ErrDBPersistenceFailure(err error) *customErrors.PublicError {
	// We use error wrapping (%w) to preserve the original GORM/driver error for logging
	return customErrors.NewHTTPError(
		http.StatusInternalServerError, 
		"Failed to save user record due to an unexpected database error.",
		fmt.Errorf("db persistence error: %w", err),
	)
}

// ErrDBLookupFailure is for errors during a SELECT/Check operation (e.g., connection issue).
// We update this slightly to accept an explicit message.
func ErrDBLookupFailure(message string, err error) *customErrors.PublicError {
	return customErrors.NewHTTPError(
		http.StatusInternalServerError, 
		message,
		fmt.Errorf("db lookup error: %w", err),
	)
}

// isUniqueConstraintViolation is a mock function. 
// In a real app, this MUST check the specific driver error code (e.g., Postgres "23505").
func isUniqueConstraintViolation(err error) bool {
	// You would inspect the error type or code here. Example:
	// var pgErr *pgconn.PgError
	// return errors.As(err, &pgErr) && pgErr.Code == "23505"
	return false // Replace with real check
}

// ErrUserNotFound returns a 404 Not Found error for when a lookup fails.
func ErrUserNotFound(lookupKey string, lookupValue interface{}) *customErrors.PublicError {
	message := fmt.Sprintf("User not found with %s: %v.", lookupKey, lookupValue)
	return customErrors.NewHTTPError(
		http.StatusNotFound, 
		message,
		nil, // gorm.ErrRecordNotFound is not wrapped, as it's a known state
	)
}