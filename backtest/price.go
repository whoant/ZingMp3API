package backtest

import (
	"time"
)

type DataPoint struct {
	Open  float64   `json:"open"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
	Close float64   `json:"close"`
	Time  time.Time `json:"timestamp"`
}

// OHLCT - Represents a datapoint in candle format that contain Open, High, Low, Close, Timestamp
type OHLCT interface {

	// Open is the starting price for a candlestick
	OpenPrice() float64

	// Close is the finish price when a candlestick that concluded
	ClosePrice() float64

	// High is the highest price reached between the Time a candlestick is Open and Close
	HighPrice() float64

	// Low is the lowest price reached between the Time a candlestick is Open and Close
	LowPrice() float64

	// Timestamp is the Time that Open candlestick
	Timestamp() time.Time
}

// Open is the starting price for a candlestick
func (candle DataPoint) OpenPrice() float64 {
	return candle.Open
}

// Close is the finish price when a candlestick that concluded
func (candle DataPoint) ClosePrice() float64 {
	return candle.Close
}

// High is the highest price reached between the Time a candlestick is Open and Close
func (candle DataPoint) HighPrice() float64 {
	return candle.High
}

// Low is the lowest price reached between the Time a candlestick is Open and Close
func (candle DataPoint) LowPrice() float64 {
	return candle.Low
}

// Timestamp is the Time that Open candlestick
func (candle DataPoint) Timestamp() time.Time {
	return candle.Time
}
