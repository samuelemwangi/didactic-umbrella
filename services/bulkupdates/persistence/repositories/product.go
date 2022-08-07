package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type ProductRepository interface {
	GetProduct(*domain.Product) error
	SaveProduct(*domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) GetProduct(product *domain.Product) error {
	result := repo.db.Where("sku = ?", product.SKU).Find(product)
	return result.Error
}

func (repo *productRepository) SaveProduct(product *domain.Product) error {
	result := repo.db.Create(product)
	return result.Error
}
