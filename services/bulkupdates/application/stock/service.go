package stock

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type StockService interface {
	SaveStock(uint, uint, int) error
}

type stockService struct {
	stockRepo repositories.StockRepository
}

func NewStockService(repos *persistence.Repositories) *stockService {
	return &stockService{
		stockRepo: repos.StockRepo,
	}
}

func (service *stockService) SaveStock(countryId uint, productId uint, quantity int) error {
	// Get stock
	stock := &domain.Stock{
		CountryID: countryId,
		ProductID: productId,
	}
	err := service.stockRepo.GetStockByProductAndCountry(stock)

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}

	// if no stock found, create new one
	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {

		stock.Count = quantity

		if stock.Count < 0 {
			stock.Count = 0
		}

		return service.stockRepo.SaveStock(stock)
	}

	// if stock found, update it
	stock.Count = stock.Count + quantity
	if stock.Count < 0 {
		stock.Count = 0
	}

	return service.stockRepo.UpdateStock(stock)
}
