package main

import (
	"fmt"
	"os"

	"coinconv/processors"

	log "github.com/sirupsen/logrus"
)

var (
	amount      string
	convertFrom string
	convertTo   string
)

// ./coinconv 123.45 USD BTC
func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Not enough arguments want 4 have: %d", len(os.Args))
	}

	amount = os.Args[1]
	convertFrom = os.Args[2]
	convertTo = os.Args[3]

	res, err := processors.Converter(amount, convertFrom, convertTo)
	if err != nil {
		log.
			WithError(err).
			Error("error while converting")
	}

	fmt.Println(res)

}
