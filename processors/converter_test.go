package processors

import (
	"coinconv/models"
	"coinconv/services/coinmarket"
	"testing"
)

func TestConverter(t *testing.T) {
	coinMarketVar := &models.CoinMarket{
		CoinMarketKey: "aac8340f-4ae2-4f14-80b0-f11fe74fde55",
		URL:           "https://pro-api.coinmarketcap.com",
		Amount:        "-10",
		ConvertFrom:   "USD",
		ConvertTo:     "EUR",
	}

	_, err := Converter(coinMarketVar)
	if err != coinmarket.APIKeyIsInvalid {
		t.Errorf("Error expect: API Key is invalid, has: %s", err)
	}
}