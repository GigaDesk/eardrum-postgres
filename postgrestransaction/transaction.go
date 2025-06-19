package postgrestransaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
    TotalAmountInCents int64
	TransactionCostInCents int64
}


// Returns the unique ID of the transaction
func (t Transaction) GetID() int64 {
	return int64(t.ID)
}

// Returns the creation timestamp of the transaction
func (t Transaction) GetCreatedAt() time.Time {
	return t.CreatedAt.UTC()
}

// Returns the update timestamp of the transaction
func (t Transaction) GetUpdatedAt() time.Time {
	return t.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the transaction
func (t Transaction) GetDeletedAt() time.Time {
	return t.DeletedAt.Time.UTC()
}

// Returns the total amount in cents spent in the transaction
func (t Transaction) GetTotalAmountInCents() int64 {
	return t.TotalAmountInCents
}

// Returns the transaction cost in cents spent in the transaction
func (t Transaction) GetTransactionCostInCents() int64 {
	return t.TransactionCostInCents
}


