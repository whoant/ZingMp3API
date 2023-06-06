package strategy

import (
	backtest2 "new-back-testing/backtest"
)

//strategy := &strategy2.AvellanedaMarketMakingStrategy{Step: 0}

type AvellanedaMarketMakingStrategy struct {
	Step int
}

func (m *AvellanedaMarketMakingStrategy) Naming() string {
	return "Avellaneda Market Making Strategy"
}

func (m *AvellanedaMarketMakingStrategy) OpenNewOrder(price backtest2.DataPoint) *backtest2.OpeningOrder {
	m.Step++
	if m.Step%2 == 0 {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.ASK,
			TakeProfitPrice:  price.OpenPrice() + 10,
			CancelOrderPrice: price.OpenPrice() - (price.OpenPrice() * 0.1),
		}
	} else {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.BID,
			TakeProfitPrice:  price.OpenPrice() - 10,
			CancelOrderPrice: price.OpenPrice() + (price.OpenPrice() * 0.1),
		}
	}
}
