package strategy

import (
	backtest2 "new-back-testing/backtest"
)

type FixedGridStrategy struct {
	Step          int
	askFirstPrice float64
	bidFirstPrice float64
}

func (m *FixedGridStrategy) Naming() string {
	return "Fixed Grid"
}

func (m *FixedGridStrategy) OpenNewOrder(price backtest2.DataPoint) *backtest2.OpeningOrder {
	if m.Step == 0 {
		m.askFirstPrice = price.OpenPrice()
		m.bidFirstPrice = price.OpenPrice()
	}

	m.Step++
	if m.Step%2 == 0 {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.ASK,
			TakeProfitPrice:  price.OpenPrice() + (price.OpenPrice() * 0.01),
			CancelOrderPrice: price.OpenPrice() - (price.OpenPrice() * 0.1),
		}
	} else {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.BID,
			TakeProfitPrice:  price.OpenPrice() - (price.OpenPrice() * 0.01),
			CancelOrderPrice: price.OpenPrice() + (price.OpenPrice() * 0.1),
		}
	}
}
