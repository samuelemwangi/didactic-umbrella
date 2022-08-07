package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/uploadmetadata"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
)

type Services struct {
	CountryService        country.CountryService
	ProductService        product.ProductService
	StockService          stock.StockService
	UploadMetadataService uploadmetadata.UploadMetadataService
	ErrorService          errorhelper.ErrorService
}

func NewServices(repos *persistence.Repositories) *Services {
	return &Services{
		CountryService:        country.NewCountryService(repos),
		ProductService:        product.NewProductService(repos),
		StockService:          stock.NewStockService(repos),
		UploadMetadataService: uploadmetadata.NewUploadMetadataService(repos),
		ErrorService:          errorhelper.NewErrorService(),
	}
}
