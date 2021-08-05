package usecases

import (
	"coinconv/domain"
)

// MarketInteractor belong to the usecases layer
type MarketInteractor struct {
	MarketRepository MarketRepository
}

// Convert ...
func (mi *MarketInteractor) Convert(input domain.ConvertInput) (res float64, err error) {
	res, err = mi.MarketRepository.Convert(input)

	return
}
