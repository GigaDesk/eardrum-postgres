package mocktransaction

import (
	"time"
)

var MultipleTransactionNodes = []MockTransaction{
	{
		Id:                      3,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      105,
		BalanceBeforeInCents: 2005,
		BalanceAfterInCents:  1900,
	},
	{
		Id:                      4,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      205,
		BalanceBeforeInCents: 2105,
		BalanceAfterInCents:  1900,
	},
	{
		Id:                      5,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      305,
		BalanceBeforeInCents: 2205,
		BalanceAfterInCents:  1900,
	},
	{
		Id:                      6,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      405,
		BalanceBeforeInCents: 2305,
		BalanceAfterInCents:  1900,
	},
}
