package postgrespurchase

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	TransactionID    int64
	ProductID        int64
	UnitsBought      int
	TotalAmountInCents int64
}

// Returns the unique ID of the purchase
func (p Purchase) GetID() int64 {
	return int64(p.ID)
}

// Returns the Product ID of the purchase
func (p Purchase) GetProductID() int64 {
	return int64(p.ProductID)
}

// Returns the Transaction ID of the purchase
func (p Purchase) GetTransactionID() int64 {
	return p.TransactionID
}

// Returns the units bought in the purchase
func (p Purchase) GetUnitsBought() int {
	return p.UnitsBought
}

// Returns the total amount in cents spent on the purchase
func (p Purchase) GetTotalAmountInCents() int64 {
	return p.TotalAmountInCents
}
