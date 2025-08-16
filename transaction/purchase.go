package transaction

import (
	"github.com/GigaDesk/eardrum-postgres/product"
	"gorm.io/gorm"
)

// Purchase represents a single item bought within a transaction.
// This is also known as a line item.
type Purchase struct {
    gorm.Model
    
    // TransactionID is the foreign key linking to the transaction it belongs to.
    // It's a `uint` to match GORM's primary key type and is marked as non-nullable.
    TransactionID    uint      `gorm:"not null"`
    
    // Transaction is the GORM association to the Transaction model.
    Transaction      Transaction `gorm:"foreignKey:TransactionID"`

    // ProductID is the foreign key linking to the specific product that was purchased.
    // It's a `uint` and is also non-nullable.
    ProductID        uint      `gorm:"not null"`
    
    // Product is the GORM association to the Product model.
    Product          product.Product `gorm:"foreignKey:ProductID"`

    // UnitsBought is the quantity of the product purchased.
    // We use `uint` as you can't buy a negative number of units and `not null` as it's required.
    UnitsBought      uint      `gorm:"not null"`
    
    // TotalAmountInCents is the total cost of this specific purchase item, in cents.
    // We use `uint` to prevent negative values and `not null` to ensure it's always present.
    TotalAmountInCents uint `gorm:"not null"`
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
	return int64(p.TransactionID)
}

// Returns the units bought in the purchase
func (p Purchase) GetUnitsBought() int {
	return int(p.UnitsBought)
}

// Returns the total amount in cents spent on the purchase
func (p Purchase) GetTotalAmountInCents() int64 {
	return int64(p.TotalAmountInCents)
}
