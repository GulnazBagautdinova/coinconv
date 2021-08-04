package models

import "coinconv/configs"

type ConversionResult struct {
	Status `json:"status"`
	Data   `json:"data"`
	//Amount       string    `json:"amount"`
}

type Status struct {
	ErrorCode    uint32 `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Data struct {
	Symbol string                            `json:"symbol"`
	Quote  map[string]map[string]interface{} `json:"quote"`
}

type CoinMarket struct {
	CoinMarketKey string
	URL           string
	Amount        string
	ConvertFrom   string
	ConvertTo     string
}

func ToCoinMarket(amount, convertFrom, convertTo string, mainConfig configs.CoinconvApiOptions) *CoinMarket {
	return &CoinMarket{
		CoinMarketKey: mainConfig.CoinMarketKey,
		URL:           mainConfig.URL,
		Amount:        amount,
		ConvertFrom:   convertFrom,
		ConvertTo:     convertTo,
	}
}
