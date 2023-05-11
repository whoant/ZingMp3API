package backtest

import (
	"time"
)

type DataPoint struct {
	open, high, low, close float64
	timestamp              time.Time
}

// OHLCT - Represents a datapoint in candle format that contain Open, High, Low, Close, Timestamp
type OHLCT interface {

	// Open is the starting price for a candlestick
	Open() float64

	// Close is the finish price when a candlestick that concluded
	Close() float64

	// High is the highest price reached between the time a candlestick is open and close
	High() float64

	// Low is the lowest price reached between the time a candlestick is open and close
	Low() float64

	// Timestamp is the time that open candlestick
	Timestamp() time.Time
}

// Open is the starting price for a candlestick
func (candle DataPoint) Open() float64 {
	return candle.open
}

// Close is the finish price when a candlestick that concluded
func (candle DataPoint) Close() float64 {
	return candle.close
}

// High is the highest price reached between the time a candlestick is open and close
func (candle DataPoint) High() float64 {
	return candle.high
}

// Low is the lowest price reached between the time a candlestick is open and close
func (candle DataPoint) Low() float64 {
	return candle.low
}

// Timestamp is the time that open candlestick
func (candle DataPoint) Timestamp() time.Time {
	return candle.timestamp
}
