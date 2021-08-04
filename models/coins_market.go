package models

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
	Symbol string                  `json:"symbol"`
	Quote  map[string]QuoteDetails `json:"quote"`
}

type QuoteDetails map[string]interface{}
