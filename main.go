package main

import (
	"flag"

	"coinconv/configs"
	"coinconv/infrastructure"

	log "github.com/sirupsen/logrus"
)

var (
	configName  string
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

	marketHandler := infrastructure.NewMarketHandler(mainConfig)

	infrastructure.Dispatch(marketHandler)
}
