package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/upload"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
)

type Services struct {
	CountryService country.CountryService
	ProductService product.ProductService
	StockService   stock.StockService
	UploadService  upload.UploadService
}

func NewServices(repos *persistence.Repositories) *Services {
	return &Services{
		CountryService: country.NewCountryService(repos),
		ProductService: product.NewProductService(repos),
		StockService:   stock.NewStockService(repos),
		UploadService:  upload.NewUploadService(repos),
	}
}
