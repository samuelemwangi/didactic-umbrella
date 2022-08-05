package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductRepository interface {
	GetProductBySKU(sku string) (*domain.Product, *string)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetProductBySKU(sku string) (*domain.Product, *string) {
	var product domain.Product

	err := pr.db.Debug().Where("sku=?", sku).Take(&product).Error
	if err != nil {
		errorMessage := err.Error()
		return nil, &errorMessage
	}

	return &product, nil
}
