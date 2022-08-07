package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductRepository interface {
	GetProductBySKU(*domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) GetProductBySKU(product *domain.Product) error {
	result := repo.db.First(product, "sku = ?", product.SKU)
	return result.Error
}
