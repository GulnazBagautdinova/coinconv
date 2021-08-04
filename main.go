package main

import (
	"flag"
	"fmt"
	"os"

	"coinconv/configs"
	"coinconv/models"
	"coinconv/processors"

	log "github.com/sirupsen/logrus"
)

var (
	configName string

	amount      string
	convertFrom string
	convertTo   string
)

func init() {
	flag.StringVar(&configName, "config", "coinconv", "configuration file name")
}

func main() {
	flag.Parse()

	// viper
	v, err := configs.NewViper(configName)
	if err != nil {
		log.Fatal(err)
	}

	// config
	mainConfig := configs.NewMainConfig(v)

	if len(os.Args) != 4 {
		log.Fatalf("Not enough arguments want 4 have: %d", len(os.Args))
	}

	amount = os.Args[1]
	convertFrom = os.Args[2]
	convertTo = os.Args[3]

	coinMarketVar := models.ToCoinMarket(amount, convertFrom, convertTo, mainConfig)

	res, err := processors.Converter(coinMarketVar)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Converting result: %f\n", res)
}
