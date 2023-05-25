package backtest

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// DataHandler is a wrapper that packages the required data for running back testing simulation.
type DataHandler struct {
	Prices []DataPoint
}

// Required columns in the CSV file
var csvColumns = []string{"timestamp", "open", "high", "low", "close"}

// newDataHandler creates and initializes a DataHandler with pricing data and executes the required setup
func newDataHandler(prices []DataPoint) *DataHandler {
	return &DataHandler{
		Prices: prices,
	}
}

// PricesFromCSV reads all csv data in the OHLCV format to the DataHandler and returns if a error occurred
func PricesFromCSV(csvFilePath string) (*DataHandler, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot OpenPrice file : %v", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Reading first line header and validating the required columns
	if line, err := reader.Read(); err != nil || !isCSVHeaderValid(line) {
		log.Println(isCSVHeaderValid(line))
		return nil, fmt.Errorf(`error reading header with columns in the csv.
				Make sure the CSV has the columns Timestamp, OpenPrice, HighPrice, LowPrice, ClosePrice`)
	}

	var prices []DataPoint
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//Checking each OHLCV value in the csv
		var numbers [5]float64

		for i := 0; i < len(numbers); i++ {
			value, err := strToFloat(line[i])
			if err != nil {
				return nil, err
			}
			numbers[i] = value

		}

		prices = append(prices, DataPoint{
			Time:  floatToTime(numbers[0]),
			Open:  numbers[1],
			High:  numbers[2],
			Low:   numbers[3],
			Close: numbers[4],
		})
	}

	return newDataHandler(prices), nil
}

// strToFloat converts a string value to float64, in case of error Panic
func strToFloat(str string) (float64, error) {
	number, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return number, nil
	}
	return -1, fmt.Errorf(`invalid parameter '%v' was found in the provided csv. 
		Make sure the csv contain only valid float numbers`, str)
}

func floatToTime(number float64) time.Time {
	sec, dec := math.Modf(number)

	return time.Unix(int64(sec), int64(dec*(1e9)))
}

// Check if the first line with columns of the csv are in the valid format
func isCSVHeaderValid(firstLine []string) bool {
	for i, column := range csvColumns {
		if strings.ToLower(firstLine[i]) != column {
			return false
		}
	}
	return true
}
