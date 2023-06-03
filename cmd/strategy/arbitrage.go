package strategy

import (
	"github.com/RobinUS2/golang-moving-average"
	backtest2 "new-back-testing/backtest"
)

//strategy := &strategy2.ArbitrageStrategy{Sma: movingaverage.New(100)}

type ArbitrageStrategy struct {
	Sma *movingaverage.MovingAverage
}

func (m *ArbitrageStrategy) Naming() string {
	return "Arbitrage Strategy"
}

func (m *ArbitrageStrategy) OpenNewOrder(price backtest2.DataPoint) *backtest2.OpeningOrder {
	m.Sma.Add(price.ClosePrice())
	if m.Sma.Count() == 100 {
		if price.OpenPrice() > m.Sma.Avg() {
			return &backtest2.OpeningOrder{
				OrderType:        backtest2.ASK,
				TakeProfitPrice:  price.OpenPrice() + (price.OpenPrice() * 0.05),
				CancelOrderPrice: price.OpenPrice() - (price.OpenPrice() * 0.1),
			}
		}

		if price.OpenPrice() < m.Sma.Avg() {
			return &backtest2.OpeningOrder{
				OrderType:        backtest2.BID,
				TakeProfitPrice:  price.OpenPrice() - (price.OpenPrice() * 0.05),
				CancelOrderPrice: price.OpenPrice() + (price.OpenPrice() * 0.1),
			}
		}
	}
	return nil
}
