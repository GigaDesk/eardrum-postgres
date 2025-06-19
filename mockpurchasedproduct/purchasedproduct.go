package mockpurchasedproduct

type Mockpurchasedproduct struct {
	ProductID   int64
	UnitsBought int
}

func (m Mockpurchasedproduct) GetProductID() int64 {
	return m.ProductID
}

func (m Mockpurchasedproduct) GetUnitsBought() int {
	return m.UnitsBought
}

