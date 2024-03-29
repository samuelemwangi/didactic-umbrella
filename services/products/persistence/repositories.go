package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type Repositories struct {
	CountryRepo        repositories.CountryRepository
	ProductRepo        repositories.ProductRepository
	StockRepo          repositories.StockRepository
	UploadMetadataRepo repositories.UploadMetadataRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		CountryRepo:        repositories.NewCountryRepository(db),
		ProductRepo:        repositories.NewProductRepository(db),
		StockRepo:          repositories.NewStockRepository(db),
		UploadMetadataRepo: repositories.NewUploadMetadataRepository(db),
	}

}
