package backtest

type Strategy interface {
	OpenNewOrder(price DataPoint) *OpeningOrder
	Naming() string
}
