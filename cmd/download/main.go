package main

import (
	"flag"
	"strings"

	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	scraper "new-back-testing/scraper"
)

var (
	from     string
	to       string
	interval string
	symbol   string
)

func main() {
	flag.StringVar(&from, "from", "", "From (MM/DD/YYYY)")
	flag.StringVar(&to, "to", "", "To (MM/DD/YYYY)")
	flag.StringVar(&interval, "interval", "", "Interval (1m, 3m, 5m, 15m, 30m, 1h)")
	flag.StringVar(&symbol, "symbol", "", "Symbol (BTC/USDT)")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	s := scraper.NewBinanceScraper()

	validFrom(from)
	validTo(to)
	validInterval(interval)
	validSymbolFormat(symbol)

	fromTime, _ := time.Parse("02/01/2006", from)
	toTime, _ := time.Parse("02/01/2006", to)
	fileName := s.Download(fromTime, toTime, interval, strings.Join(strings.Split(symbol, "/"), ""))
	log.Info().Msgf("Download successful !! Save file : %v", fileName)
}

func validFrom(t string) {
	fromTime, err := time.Parse("02/01/2006", t)
	if err != nil {
		log.Fatal().Err(err).Msg("From is invalid")
	}

	if fromTime.Before(time.Now()) {
		log.Info().Msg("From is valid")
		return
	}

	log.Fatal().Msg("From is invalid")
}

func validTo(t string) {
	fromTime, err := time.Parse("02/01/2006", t)
	if err != nil {
		log.Fatal().Err(err).Msg("To is invalid")
	}

	if fromTime.Before(time.Now()) {
		log.Info().Msg("To is valid")
		return
	}

	log.Fatal().Msg("To is invalid")
}

func validInterval(interval string) {
	switch interval {
	case "1m", "3m", "5m", "15m", "30m", "1h":
		log.Info().Msg("Interval is valid")
		return
	}

	log.Fatal().Msg("Interval is invalid")
}

func validSymbolFormat(symbol string) {
	if len(strings.Split(symbol, "/")) == 2 {
		log.Info().Msg("Format of symbol is valid")
		return
	}
	log.Fatal().Msg("Format of symbol is invalid")
}
