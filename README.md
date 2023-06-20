# BackTesting Core

## Feature
- Download backtesting data from Binance
- Backtest strategy is a simple
- Statistic profit, profit margin, and CAGR after backtesting
- UI shows orders history in the table and candlestick chart

## Download data
1. Enter in terminal ```go build -o out/download cmd/download/main.go``` to build download tool
2. Next, enter in terminal ``` ./out/download --from 01/06/2023 --to 10/06/2023 --interval 1h --symbol BTC/USDT  ``` to download data from 01/06/2023 to 10/06/2023, the interval is 1h and symbol is BTC/USDT. The file will be exported to the main dir.
   
## Backtesting
1. You can see an example file in ``` cmd/backtest/main.go ``` and define strategy in ``` cmd/strategy ``` folder
