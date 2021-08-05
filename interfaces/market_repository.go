package interfaces

import (
	"coinconv/domain"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
)

// MarketRepository belong to the inteface layer
type MarketRepository struct {
	MarketHandler MarketHandler
}

// Convert ...
func (mr *MarketRepository) Convert(arg domain.ConvertInput) (res float64, err error) {
	resp, err := mr.MarketHandler.DoRequest(arg.Amount, arg.ConvertFrom, arg.ConvertTo)
	if err != nil {
		log.Error(err)
		return
	}

	switch resp.StatusCode {
	case 200, 400:
	case 401:
		err = domain.APIKeyIsInvalid
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
	signDeteils := quoteDetails[strings.ToUpper(arg.ConvertTo)].(map[string]interface{})
	res = signDeteils["price"].(float64)
	return
}
