package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type StockRepository interface {
	GetStockByProductAndCountry(*domain.Stock) (*domain.Stock, *string)
	UpdateStockCount(*domain.Stock) (*domain.Stock, *string)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{db}
}

func (sr *stockRepository) GetStockByProductAndCountry(stockTosearch *domain.Stock) (*domain.Stock, *string) {
	var stock domain.Stock

	err := sr.db.Debug().Where(stockTosearch).Take(&stock).Error

	if err != nil {
		errorMessage := err.Error()
		return nil, &errorMessage
	}

	return &stock, nil
}

func (sr *stockRepository) UpdateStockCount(stockToUpdate *domain.Stock) (*domain.Stock, *string) {
	result := sr.db.Model(&domain.Stock{}).Updates(stockToUpdate)
	if result.Error != nil {
		errorMessage := result.Error.Error()
		return nil, &errorMessage
	}
	return stockToUpdate, nil
}
