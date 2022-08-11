package stock

import (
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

	stock, err := service.stockRepo.GetStockByProductAndCountry(countryId, productId)

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}

	// if no stock found, create new one
	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {

		stock.CountryID = countryId
		stock.ProductID = productId
		stock.Quantity = quantity

		if stock.Quantity < 0 {
			stock.Quantity = 0
		}

		return service.stockRepo.SaveStock(stock)
	}

	// if stock found, update it
	stock.Quantity = stock.Quantity + quantity
	if stock.Quantity < 0 {
		stock.Quantity = 0
	}

	return service.stockRepo.UpdateStock(stock)
}
