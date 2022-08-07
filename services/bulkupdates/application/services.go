package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/uploadmetadata"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
)

type Services struct {
	CountryService        country.CountryService
	ProductService        product.ProductService
	StockService          stock.StockService
	UploadMetadataService uploadmetadata.UploadMetadataService
}

func NewServices(repos *persistence.Repositories) *Services {
	return &Services{
		CountryService:        country.NewCountryService(repos),
		ProductService:        product.NewProductService(repos),
		StockService:          stock.NewStockService(repos),
		UploadMetadataService: uploadmetadata.NewUploadMetadataService(repos),
	}
}
