package main

import (
	"github.com/rs/zerolog/log"

	backtest2 "new-back-testing/backtest"
)

type myStrategy struct {
	step int
}

func (m *myStrategy) Naming() string {
	return "Market marker"
}

func (m *myStrategy) OpenNewOrder(price backtest2.DataPoint) *backtest2.OpeningOrder {
	m.step++
	if m.step%2 == 0 {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.ASK,
			TakeProfitPrice:  price.Open() + (price.Open() * 0.01),
			CancelOrderPrice: price.Open() - (price.Open() * 0.1),
		}
	} else {
		return &backtest2.OpeningOrder{
			OrderType:        backtest2.BID,
			TakeProfitPrice:  price.Open() + (price.Open() * 0.01),
			CancelOrderPrice: price.Open() - (price.Open() * 0.1),
		}
	}

}

func main() {
	dataHandler, err := backtest2.PricesFromCSV("./BTCUSDT|1h|09-04-2023 00:00|11-04-2023 00:00.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load data")
	}

	strategy := &myStrategy{step: 0}

	backTestOptions := backtest2.NewBacktestOptions("BTC/USDT", 10, 100, 100)

	a := backtest2.NewBackTest(strategy, dataHandler, backTestOptions)
	a.Run()
	a.Portfolio()
}
