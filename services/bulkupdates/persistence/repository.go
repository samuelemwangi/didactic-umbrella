package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
)

type Repositories struct {
	CountryRepo       repositories.CountryRepository
	ProductRepo       repositories.ProductRepository
	StockRepo         repositories.StockRepository
	UploadMetdataRepo repositories.UploadMetadataRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		CountryRepo:       repositories.NewCountryRepository(db),
		ProductRepo:       repositories.NewProductRepository(db),
		StockRepo:         repositories.NewStockRepository(db),
		UploadMetdataRepo: repositories.NewUploadMetadataRepository(db),
	}

}
