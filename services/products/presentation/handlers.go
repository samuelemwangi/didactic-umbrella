package presentation

import (
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation/handlers"
)

type Handlers struct {
	CountryHandler handlers.CountryHandler
	ProductHandler handlers.ProductHandler
	StockHandler   handlers.StockHandler
}

func NewHandlers(services *application.Services) *Handlers {
	return &Handlers{
		CountryHandler: *handlers.NewCountryHandler(services),
		ProductHandler: *handlers.NewProductHandler(services),
		StockHandler:   *handlers.NewStockHandler(services),
	}
}
