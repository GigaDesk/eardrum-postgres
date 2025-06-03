package mockproduct

import (
	"time"
)

var MultipleProductNodes = []MockProduct{
	{
		Id:                      3,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Pen",
		Price_per_unit_in_cents: 5600,
	},
	{
		Id:                      4,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Biscuit",
		Price_per_unit_in_cents: 7000,
	},
	{
		Id:                      5,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Soda",
		Price_per_unit_in_cents: 8000,
	},
	{
		Id:                      6,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Lotion",
		Price_per_unit_in_cents: 6000,
	},
}

var UpdatedProduct = MockProduct{
	Id:                      3,
	CreatedAt:               time.Now().UTC(),
	UpdatedAt:               time.Now().UTC(),
	Name:                    "A4 Book",
	Price_per_unit_in_cents: 2000,
}
