package interfaces

import (
	"coinconv/domain"
	"coinconv/usecases"
)

// MarketController belong to the interface layer
type MarketController struct {
	MarketInteractor usecases.MarketInteractor
	//Logger         usecases.Logger
}

// NewMarketController ...
func NewMarketController(marketHandler MarketHandler) *MarketController {
	return &MarketController{
		MarketInteractor: usecases.MarketInteractor{
			MarketRepository: &MarketRepository{
				MarketHandler: marketHandler,
			},
		},
		// TODO: add logger
		//Logger: logger,
	}
}

// DoConvertion ...
func (mc *MarketController) DoConvertion(arg domain.ConvertInput) (res float64, err error) {
	res, err = mc.MarketInteractor.Convert(arg)
	return
}
