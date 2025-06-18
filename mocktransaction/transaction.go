package mocktransaction

import (
	"time"

	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/GigaDesk/eardrum-postgres/mockpurchasedproduct"
)

type MockTransaction struct {
	Id                   int64
	CreatedAt            time.Time
	UpdatedAt            time.Time
	TotalAmountInCents   int64
	BalanceBeforeInCents int64
	BalanceAfterInCents  int64
}

func (m MockTransaction) GetID() int64 {
	return m.Id
}
func (m MockTransaction) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockTransaction) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockTransaction) GetDeletedAt() time.Time {
	return time.Time{}
}
func (m MockTransaction) GetTotalAmountInCents() int64 {
	return m.TotalAmountInCents
}
func (m MockTransaction) GetBalanceBeforeInCents() int64 {
	return m.BalanceBeforeInCents
}
func (m MockTransaction) GetBalanceAfterInCents() int64 {
	return m.BalanceAfterInCents
}

type MockNewTransaction struct {
	PurchasedProducts  []mockpurchasedproduct.Mockpurchasedproduct
	PhoneNumber string
	PINCode            string
}

func (m MockNewTransaction) GetPurchasedProducts() []transaction.PurchasedProduct {

	var purchasedproducts []transaction.PurchasedProduct

	for _, p := range m.PurchasedProducts {
		purchasedproducts = append(purchasedproducts, p)
	}

	return purchasedproducts
}

func (m MockNewTransaction) GetPhoneNumber() string {

	return m.PhoneNumber
}

func (m MockNewTransaction) GetPinCode() string {

	return m.PINCode
}

