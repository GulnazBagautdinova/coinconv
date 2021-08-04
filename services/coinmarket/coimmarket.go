package coinmarket

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"coinconv/models"
)

const conversionURL = "/v1/tools/price-conversion"

var (
	APIKeyIsInvalid = errors.New("API Key is invalid")
)

func CoinmarketAPICall(coinMarketVar *models.CoinMarket) (respBody []byte, err error) {
	client := &http.Client{}
	urlMarket := coinMarketVar.URL + conversionURL
	req, err := http.NewRequest("GET", urlMarket, nil)
	if err != nil {
		log.Print(err)
		return
	}

	q := url.Values{}
	q.Add("amount", coinMarketVar.Amount)
	q.Add("symbol", coinMarketVar.ConvertFrom)
	q.Add("convert", coinMarketVar.ConvertTo)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinMarketVar.CoinMarketKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case 200, 400:
	case 401:
		err = APIKeyIsInvalid
		return
	default:
		err = errors.New(resp.Status)
		return
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
