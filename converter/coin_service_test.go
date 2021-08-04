package converter

import (
	"coinconv/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mainConfig = configs.CoinconvApiOptions{
		CoinMarketKey: "aac8340f-4ae2-4f14-80b0-f11fe74fde55",
		URL:           "https://pro-api.coinmarketcap.com",
	}
)

func TestConvertWrongKey(t *testing.T) {
	var (
		amount      = "10"
		convertFrom = "USD"
		convertTo   = "EUR"
	)

	coinMarketService := NewCoinMarketService(mainConfig)

	_, err := coinMarketService.Convert(amount, convertFrom, convertTo)
	assert.NotNil(t, err)

}
