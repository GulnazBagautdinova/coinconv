package coinmarket

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"coinconv/configs"
	"coinconv/services/interfaces/converter"
)

const conversionURL = "/v1/tools/price-conversion"

var (
	APIKeyIsInvalid = errors.New("API Key is invalid")
)

// coinMarketService presents a connection to CoinMarketCap
type coinMarketService struct {
	httpClient *http.Client
	apiKey     string
	apiURL     string
}

// NewCoinMarketService return Coin Market service which realize Service interface
func NewCoinMarketService(opt configs.CoinconvApiOptions) converter.Service {
	return &coinMarketService{
		httpClient: &http.Client{},
		apiURL:     opt.URL,
		apiKey:     opt.CoinMarketKey,
	}
}

// Convert coins using https://coinmarketcap.com/api
func (cm *coinMarketService) Convert(amount, convertFrom, convertTo string) (res float64, err error) {
	urlMarket := cm.apiURL + conversionURL
	req, err := http.NewRequest("GET", urlMarket, nil)
	if err != nil {
		log.Print(err)
		return
	}

	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", convertFrom)
	q.Add("convert", convertTo)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cm.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := cm.httpClient.Do(req)
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

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	data := make(map[string]map[string]interface{})
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		err = errors.New("error reading json")
		return
	}

	if status, ok := data["status"]; ok && status["error_message"] != nil {
		err = errors.New(status["error_message"].(string))
		return
	}

	quoteDetails := data["data"]["quote"].(map[string]interface{})
	signDeteils := quoteDetails[strings.ToUpper(convertTo)].(map[string]interface{})
	res = signDeteils["price"].(float64)

	return
}
