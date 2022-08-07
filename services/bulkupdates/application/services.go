package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/uploadprocess"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
)

type Services struct {
	CountryService         country.CountryService
	ProductService         product.ProductService
	StockService           stock.StockService
	UploadProcessorService uploadprocess.UploadProcessorService
}

func NewServices(repos *persistence.Repositories) *Services {
	return &Services{
		CountryService:         country.NewCountryService(repos),
		ProductService:         product.NewProductService(repos),
		StockService:           stock.NewStockService(repos),
		UploadProcessorService: uploadprocess.NewUploadProcessorService(repos),
	}
}
