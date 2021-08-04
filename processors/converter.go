package processors

import (
	"encoding/json"
	"errors"
	"strings"

	"coinconv/models"
	"coinconv/services/coinmarket"

	log "github.com/sirupsen/logrus"
)

func Converter(coinMarketVar *models.CoinMarket) (res float64, err error) {
	respBody, err := coinmarket.CoinmarketAPICall(coinMarketVar)
	if err != nil {
		return
	}

	var v models.ConversionResult
	err = json.Unmarshal(respBody, &v)
	if err != nil {
		log.
			WithError(err).
			Error("error reading json")
		return
	}

	if v.ErrorCode != 0 {
		err = errors.New(v.ErrorMessage)
		return
	}

	quoteDetails := v.Quote[strings.ToUpper(coinMarketVar.ConvertTo)]
	res = quoteDetails["price"].(float64)
	return
}
