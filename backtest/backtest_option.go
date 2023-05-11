package backtest

import (
	"strings"

	"github.com/rs/zerolog/log"
)

type BacktestOptions struct {
	Pair               string // BTC/USDT
	AmountPerOrder     float64
	InitialQuoteAmount float64 // USDT
	InitialBaseAmount  float64 // BTC
	CurrentQuoteAmount float64
	CurrentBaseAmount  float64
}

func NewBacktestOptions(pair string, amountPerOrder float64, initialQuoteAmount float64, initialBaseAmount float64) *BacktestOptions {
	return &BacktestOptions{
		Pair:               pair,
		AmountPerOrder:     amountPerOrder,
		InitialBaseAmount:  initialBaseAmount,
		CurrentBaseAmount:  initialBaseAmount,
		InitialQuoteAmount: initialQuoteAmount,
		CurrentQuoteAmount: initialQuoteAmount,
	}
}

func (options *BacktestOptions) Portfolio() {
	coins := strings.Split(options.Pair, "/")
	baseCoin, quoteCoin := coins[0], coins[1]
	log.Info().Str("pair", options.Pair).
		Str("base_coin", baseCoin).
		Str("quote_coin", quoteCoin).
		Float64("amount_per_order", options.AmountPerOrder).
		Float64("initial_base_amount", options.InitialBaseAmount).
		Float64("current_base_amount", options.CurrentBaseAmount).
		Float64("initial_quote_amount", options.InitialQuoteAmount).
		Float64("current_quote_amount", options.CurrentQuoteAmount).
		Msg("portfolio")
}
