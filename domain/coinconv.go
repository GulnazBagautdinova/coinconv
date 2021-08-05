package domain

type ConvertInput struct {
	Amount      string
	ConvertFrom string
	ConvertTo   string
}

// func ToCoinMarket(amount, convertFrom, convertTo string, mainConfig configs.CoinconvApiOptions) *CoinMarket {
// 	return &CoinMarket{
// 		CoinMarketKey: mainConfig.CoinMarketKey,
// 		URL:           mainConfig.URL,
// 		Amount:        amount,
// 		ConvertFrom:   convertFrom,
// 		ConvertTo:     convertTo,
// 	}
// }
