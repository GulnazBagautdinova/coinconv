package models

// {"status":{"timestamp":"2021-08-03T15:49:27.096Z","error_code":0,"error_message":null,"elapsed":19,"credit_count":1,"notice":null},
// "data":{"id":2781,"symbol":"USD","name":"United States Dollar",
// "amount":1,"last_updated":"2021-08-03T15:48:52.000Z","quote":{"RUB":{"price":73.01289999999364,"last_updated":"2021-08-03T15:48:52.000Z"}}}}

type ConversionResult struct {
	Status `json:"status"`
	Data   `json:"data"`
	//Amount       string    `json:"amount"`
}

type Status struct {
	ErrorCode    uint32 `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message"`
}

type Data struct {
	Symbol string                            `json:"symbol,omitempty"`
	Quote  map[string]map[string]interface{} `json:"quote"`
}

// type Quote struct {
// 	//ID  uint32 `json:"id,omitempty"`
// 	Price `json:"-"`
// }

// type Price struct {
// 	//ID  uint32 `json:"id,omitempty"`
// 	PriceSum float64 `json:"price"`
// }
