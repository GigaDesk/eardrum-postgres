package postgrespurchase

import (
	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/postgresproduct"
	"gorm.io/gorm"
)

// Calculates the amount spent in cents in a particular purchase
func CalculatePurchaseAmountInCents(p transaction.PurchasedProduct, Db *gorm.DB) (int, error) {

	product, err := postgresproduct.GetProductWithId(Db, int(p.GetProductID()))

	if err != nil {
		return 0, err
	}

	TotalAmountCents := p.GetUnitsBought() * int(product.GetPricePerUnitInCents())

	return TotalAmountCents, nil

}

// creates a purchase record
func CreatePurchase(t transaction.Transaction, p transaction.PurchasedProduct, Db *gorm.DB) (purchase.Purchase, error) {

	Amount, err := CalculatePurchaseAmountInCents(p, Db)

	if err != nil {
		return nil, err
	}

	//create purchase data
	purchase := &Purchase{
		TransactionID:    t.GetID(),
		ProductID:        p.GetProductID(),
		UnitsBought:      p.GetUnitsBought(),
		TotalAmountInCents: int64(Amount),
	}

	//create a purchase record in the database and return if operation succeeds
	if err := Db.Create(purchase).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}
