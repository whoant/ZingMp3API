package backtest

type OpeningOrder struct {
	OrderType        OrderType
	TakeProfitPrice  float64
	CancelOrderPrice float64
}

type Strategy interface {
	OpenNewOrder(price DataPoint) *OpeningOrder
	Naming() string
}
