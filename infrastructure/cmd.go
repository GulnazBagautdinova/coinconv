package infrastructure

import (
	"fmt"
	"os"

	"coinconv/domain"
	"coinconv/interfaces"

	log "github.com/sirupsen/logrus"
)

// Dispatch is handle routing
func Dispatch(mh interfaces.MarketHandler) {
	if len(os.Args) != 4 {
		log.Infof("Not enough arguments for cmd interface, want 4 have: %d", len(os.Args))
		os.Exit(0)
	}

	input := domain.ConvertInput{
		Amount:      os.Args[1],
		ConvertFrom: os.Args[2],
		ConvertTo:   os.Args[3],
	}

	marketController := interfaces.NewMarketController(mh)

	res, err := marketController.DoConvertion(input)
	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("Converting result: %f\n", res)
	}
}
