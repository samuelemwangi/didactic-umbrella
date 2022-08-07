package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductRepository interface {
	GetProductBySKU(*domain.Product) error
	GetProducts() ([]*domain.Product, error)
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

func (repo *productRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	result := repo.db.Find(&products)
	return products, result.Error
}
