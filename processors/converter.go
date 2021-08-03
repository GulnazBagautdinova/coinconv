package processors

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"coinconv/models"

	log "github.com/sirupsen/logrus"
)

var coinmarketcap_url = "https://pro-api.coinmarketcap.com/v1/tools/price-conversion"

func Converter(amount, convertFrom, convertTo string) (res float64, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", coinmarketcap_url, nil)
	if err != nil {
		log.Print(err)
		return
	}

	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", convertFrom)
	q.Add("convert", convertTo)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "aac8340f-4ae2-4f14-80b0-f11fe74fde51")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		err = errors.New("Error sending request to server")
		return
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

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
	}

	res = v.Quote[convertTo]["price"].(float64)

	return
}
