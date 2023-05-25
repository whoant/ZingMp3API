package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"new-back-testing/utils"
)

const FormatFullTime = "02-01-2006 15:04"

type BinanceScraper struct {
}

func NewBinanceScraper() Scraper {
	return &BinanceScraper{}
}

func (scraper *BinanceScraper) GetData(from time.Time, to time.Time, interval string, symbol string) []OHLCT {
	var result []OHLCT
	next := to
	for next.After(from) {
		klines, err := scraper.getKLine(next, interval, symbol)
		if err != nil {
			return result
		}

		result = append(klines, result...)
		next = klines[0].Timestamp.Add(-1 * time.Millisecond)
		//log.Println(next)
		log.Info().Time("next_cursor", next).Msg("next cursor")
	}

	_, err := json.Marshal(result)
	if err != nil {
		return nil
	}

	return result
}

func (scraper *BinanceScraper) Download(from time.Time, to time.Time, interval string, symbol string) {
	data := scraper.GetData(from, to, interval, symbol)
	headers := []string{"timestamp", "open", "high", "low", "close"}
	filename := fmt.Sprintf("%v|%v|%v|%v.csv", symbol, interval, from.Format(FormatFullTime), to.Format(FormatFullTime))
	structToCsv(data, headers, filename)
}

func (scraper *BinanceScraper) getKLine(endTime time.Time, interval string, symbol string) ([]OHLCT, error) {
	var result []OHLCT

	fullUrl := fmt.Sprintf("https://www.binance.com/api/v3/uiKlines?endTime=%v&limit=100&symbol=%v&interval=%v", time2Int64(endTime), symbol, interval)
	resp, _, err := utils.HttpGet(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	var val [][]interface{}
	if err := json.Unmarshal([]byte(resp), &val); err != nil {
		return nil, err
	}

	for _, line := range val {
		open, _ := strconv.ParseFloat(line[1].(string), 32)
		high, _ := strconv.ParseFloat(line[2].(string), 32)
		low, _ := strconv.ParseFloat(line[3].(string), 32)
		close, _ := strconv.ParseFloat(line[4].(string), 32)
		newKLine := OHLCT{
			Timestamp: time.Unix(int64(line[0].(float64))/1e3, 0),
			Open:      open,
			High:      high,
			Low:       low,
			Close:     close,
		}
		result = append(result, newKLine)
	}

	return result, nil
}

func time2Int64(time time.Time) int64 {
	return time.UnixNano() / 1e6
}

func structToCsv(data []OHLCT, headers []string, filename string) error {
	var rows [][]string

	// OpenPrice the file for writing.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer.
	w := csv.NewWriter(file)

	// Write the headers to the CSV file.
	if err := w.Write(headers); err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		var row []string
		row = append(row, fmt.Sprintf("%v", data[i].Timestamp.Unix()))
		row = append(row, fmt.Sprintf("%v", data[i].Open))
		row = append(row, fmt.Sprintf("%v", data[i].High))
		row = append(row, fmt.Sprintf("%v", data[i].Low))
		row = append(row, fmt.Sprintf("%v", data[i].Close))
		rows = append(rows, row)
	}

	// Write the rows to the CSV file.
	if err := w.WriteAll(rows); err != nil {
		return err
	}

	w.Flush()

	return nil
}
