package postgrestransaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
    TotalAmountInCents int64
	BalanceBeforeInCents int64
	BalanceAfterInCents int64
}


// Returns the unique ID of the transaction
func (t Transaction) GetID() int64 {
	return int64(t.ID)
}

// Returns the creation timestamp of the transaction
func (t Transaction) GetCreatedAt() time.Time {
	return t.CreatedAt
}

// Returns the update timestamp of the transaction
func (t Transaction) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

// Returns the deletion timestamp of the transaction
func (t Transaction) GetDeletedAt() time.Time {
	return t.DeletedAt.Time
}

// Returns the total amount in cents spent in the transaction
func (t Transaction) GetTotalAmountInCents() int64 {
	return t.TotalAmountInCents
}

// Returns the balance before the transaction in cents
func (t Transaction) GetBalanceBeforeInCents() int64 {
	return t.BalanceBeforeInCents
}

// Returns the balance after the transaction in cents
func (t Transaction) GetBalanceAfterInCents() int64 {
	return t.BalanceAfterInCents
}


