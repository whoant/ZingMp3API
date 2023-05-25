package main

import (
	"time"

	scraper "new-back-testing/scraper"
)

func main() {
	s := scraper.NewBinanceScraper()
	from := "21/05/2023"
	to := "23/05/2023"

	fromTime, _ := time.Parse("02/01/2006", from)
	toTime, _ := time.Parse("02/01/2006", to)
	s.Download(fromTime, toTime, "1m", "BTCUSDT")

}
