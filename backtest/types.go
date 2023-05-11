package backtest

import (
	"time"
)

type Portfolio struct {
	Pair               string    `json:"pair"`
	BaseCoin           string    `json:"baseCoin"`
	QuoteCoin          string    `json:"quoteCoin"`
	AmountPerOrder     float64   `json:"amountPerOrder"`
	InitialBaseAmount  float64   `json:"initialBaseAmount"`
	CurrentBaseAmount  float64   `json:"currentBaseAmount"`
	InitialQuoteAmount float64   `json:"initialQuoteAmount"`
	CurrentQuoteAmount float64   `json:"currentQuoteAmount"`
	InitialSumAmount   float64   `json:"initialSumAmount"`
	CurrentSumAmount   float64   `json:"currentSumAmount"`
	Profit             float64   `json:"profit"`
	ProfitMargin       float64   `json:"profitMargin"`
	Cagr               float64   `json:"cagr"`
	Orders             []*Order  `json:"orders"`
	CreatedAt          time.Time `json:"createdAt"`
}
