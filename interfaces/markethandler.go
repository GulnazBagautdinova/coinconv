package interfaces

import "net/http"

// MarketHandler interface for working with coin convertion
type MarketHandler interface {
	DoRequest(arg1, arg2, arg3 string) (*http.Response, error)
}
