package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type StockRepository interface {
	GetStock(*domain.Stock) error
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

func (repo *stockRepository) GetStock(stock *domain.Stock) error {
	result := repo.db.Where(stock).Find(stock)
	return result.Error
}

func (repo *stockRepository) SaveStock(stock *domain.Stock) error {
	result := repo.db.Create(stock)
	return result.Error
}

func (repo *stockRepository) UpdateStock(stock *domain.Stock) error {
	result := repo.db.Model(stock).Updates(stock)
	return result.Error
}
