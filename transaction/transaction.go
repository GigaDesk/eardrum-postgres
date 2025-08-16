package transaction

import (
	"time"

	"github.com/GigaDesk/eardrum-postgres/shop"
	"github.com/GigaDesk/eardrum-postgres/user"
	"gorm.io/gorm"
)

// Transaction represents a financial transaction.
// It includes foreign keys to both a User and a Shop.
// gorm.Model automatically includes an 'ID', 'CreatedAt', 'UpdatedAt', and 'DeletedAt' fields.
type Transaction struct {
    gorm.Model
    
    // TotalAmountInCents is the transaction amount, stored in cents.
    // We use `uint` to prevent negative values and `gorm:"not null"` to ensure it's always present.
    TotalAmountInCents uint `gorm:"not null"`
    
    // TransactionCostInCents is the cost of the transaction.
    // We use `uint` and `gorm:"not null"` to enforce it's a non-negative, required field.
    TransactionCostInCents uint `gorm:"not null"`
    
    // UserID is the foreign key. The `gorm:"not null"` tag ensures
    // a transaction cannot be created without a valid user.
    UserID      uint      `gorm:"not null"`
    
    // User is the GORM association to the User model, enabling database joins.
    // The `foreignKey:UserID` tag explicitly links it to the UserID field.
    User        user.User      `gorm:"foreignKey:UserID"`
    
    // ShopID is another foreign key to the Shop model.
    ShopID      uint      `gorm:"not null"`
    
    // Shop is the GORM association.
    Shop        shop.Shop      `gorm:"foreignKey:ShopID"`
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
	return int64(t.TotalAmountInCents)
}

// Returns the transaction cost in cents spent in the transaction
func (t Transaction) GetTransactionCostInCents() int64 {
	return int64(t.TransactionCostInCents)
}

// Returns the unique identifier of the shop the transaction was made to. 🏪
func (t Transaction) GetShopID() int64 {
	return int64(t.ShopID)
}


// Returns the unique identifier of the user that made the transtion
func (t Transaction) GetUserID() int64 {
	return int64(t.UserID)
}
