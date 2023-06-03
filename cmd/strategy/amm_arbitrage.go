package strategy

import (
	"github.com/MicahParks/go-rsi/v2"
	movingaverage "github.com/RobinUS2/golang-moving-average"
	backtest2 "new-back-testing/backtest"
)

type AmmArbitrageStrategy struct {
	Sma *movingaverage.MovingAverage
	Rsi rsi.RSI
}

func (m *AmmArbitrageStrategy) Naming() string {
	return "AMM Arbitrage Strategy"
}

func (m *AmmArbitrageStrategy) OpenNewOrder(price backtest2.DataPoint) *backtest2.OpeningOrder {
	m.Sma.Add(price.ClosePrice())
	m.Rsi.Calculate(price.ClosePrice())
	if m.Sma.Count() == 100 {

		if price.OpenPrice() > m.Sma.Avg() && m.Rsi.Calculate(price.ClosePrice()) < 70 {
			return &backtest2.OpeningOrder{
				OrderType:        backtest2.ASK,
				TakeProfitPrice:  price.OpenPrice() + (price.OpenPrice() * 0.05),
				CancelOrderPrice: price.OpenPrice() - (price.OpenPrice() * 0.1),
			}
		}

		if price.OpenPrice() < m.Sma.Avg() && m.Rsi.Calculate(price.ClosePrice()) > 30 {
			return &backtest2.OpeningOrder{
				OrderType:        backtest2.BID,
				TakeProfitPrice:  price.OpenPrice() - (price.OpenPrice() * 0.05),
				CancelOrderPrice: price.OpenPrice() + (price.OpenPrice() * 0.1),
			}
		}
	}
	return nil
}
