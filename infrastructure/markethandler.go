package infrastructure

import (
	"net/http"
	"net/url"

	"coinconv/configs"
	"coinconv/interfaces"

	log "github.com/sirupsen/logrus"
)

const conversionURL = "/v1/tools/price-conversion"

// marketHandler presents a connection to CoinMarketCap
type marketHandler struct {
	httpClient *http.Client
	apiKey     string
	apiURL     string
}

// NewMarketHandler return Coin Market service which realize Service interface
func NewMarketHandler(opt configs.CoinconvApiOptions) interfaces.MarketHandler {
	return &marketHandler{
		httpClient: &http.Client{},
		apiURL:     opt.URL,
		apiKey:     opt.CoinMarketKey,
	}
}

// GetConverted coins using https://coinmarketcap.com/api
func (cm *marketHandler) DoRequest(amount, convertFrom, convertTo string) (resp *http.Response, err error) {
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
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cm.apiKey)

	resp, err = cm.httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return
	}

	// switch resp.StatusCode {
	// case 200, 400:
	// case 401:
	// 	err = APIKeyIsInvalid
	// 	return
	// default:
	// 	err = errors.New(resp.Status)
	// 	return
	// }

	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	// data := make(map[string]map[string]interface{})
	// err = json.Unmarshal(respBody, &data)
	// if err != nil {
	// 	err = errors.New("error reading json")
	// 	return
	// }

	// if status, ok := data["status"]; ok && status["error_message"] != nil {
	// 	err = errors.New(status["error_message"].(string))
	// 	return
	// }

	// quoteDetails := data["data"]["quote"].(map[string]interface{})
	// signDeteils := quoteDetails[strings.ToUpper(convertTo)].(map[string]interface{})
	// res = signDeteils["price"].(float64)

	return
}
