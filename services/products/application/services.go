package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/upload"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
)

type Services struct {
	CountryService country.CountryService
	ProductService product.ProductService
	StockService   stock.StockService
	UploadService  upload.UploadService
	ErrorService   error.ErrorService
}

func NewServices(repos *persistence.Repositories) *Services {
	return &Services{
		CountryService: country.NewCountryService(repos),
		ProductService: product.NewProductService(repos),
		StockService:   stock.NewStockService(repos),
		UploadService:  upload.NewUploadService(repos),
		ErrorService:   error.NewErrorService(),
	}
}
