package scraper

import (
	"time"
)

type OHLCT struct {
	Open      float64   `json:"open"`
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Timestamp time.Time `json:"timestamp"`
}

type Scraper interface {
	GetData(from time.Time, to time.Time, interval string, symbol string) []OHLCT
	Download(from time.Time, to time.Time, interval string, symbol string) string
}
