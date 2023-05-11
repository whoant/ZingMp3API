package main

import (
	"time"

	scraper "new-back-testing/scraper"
)

func main() {
	s := scraper.NewBinanceScraper()
	from := "08/04/2023"
	to := "11/04/2023"

	fromTime, _ := time.Parse("02/01/2006", from)
	toTime, _ := time.Parse("02/01/2006", to)
	s.Download(fromTime, toTime, "1h", "BTCUSDT")

}
