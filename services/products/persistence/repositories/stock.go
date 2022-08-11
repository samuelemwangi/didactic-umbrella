package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type StockRepository interface {
	GetStockByProductAndCountry(uint, uint) (*domain.Stock, error)
	SaveStock(*domain.Stock) error
	UpdateStock(*domain.Stock) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{
		db: db,
	}
}

func (repo *stockRepository) GetStockByProductAndCountry(countryId uint, productId uint) (*domain.Stock, error) {
	stock := &domain.Stock{}
	result := repo.db.Where("country_id = ? AND product_id = ?", stock.CountryID, stock.ProductID).Take(stock)
	return stock, result.Error
}

func (repo *stockRepository) SaveStock(stock *domain.Stock) error {
	result := repo.db.Create(stock)
	return result.Error
}

func (repo *stockRepository) UpdateStock(stock *domain.Stock) error {
	itemsToUpdate := map[string]interface{}{
		"quantity": stock.Quantity,
	}
	result := repo.db.Model(&domain.Stock{}).Where("country_id = ? AND product_id = ?", stock.CountryID, stock.ProductID).Updates(itemsToUpdate)
	return result.Error
}
