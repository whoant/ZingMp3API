package main

import (
	"github.com/rs/zerolog/log"
	strategy2 "new-back-testing/cmd/strategy"

	backtest2 "new-back-testing/backtest"
)

func main() {
	dataHandler, err := backtest2.PricesFromCSV("./BTCUSDT|1h|01-06-2023 00:00|10-06-2023 00:00.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load data")
	}

	strategy := &strategy2.AvellanedaMarketMakingStrategy{Step: 0}

	backTestOptions := backtest2.NewBacktestOptions("BTC/USDT", 10, 100000, 100)

	a := backtest2.NewBackTest(strategy, dataHandler, backTestOptions)
	a.Run()
	a.Portfolio()
}
