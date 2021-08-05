package usecases

import (
	"coinconv/domain"
)

// A MarketRepository belong to the usecases layer.
type MarketRepository interface {
	Convert(users domain.ConvertInput) (res float64, err error)
}
