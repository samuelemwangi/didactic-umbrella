package stock

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type StockService interface {
	SaveStock(*StockRequestDTO) error
}

type stockService struct {
	stockRepo repositories.StockRepository
}

func NewStockService(repos *persistence.Repositories) *stockService {
	return &stockService{
		stockRepo: repos.StockRepo,
	}
}

func (service *stockService) SaveStock(request *StockRequestDTO) error {
	// Get stock
	stock := request.toEntity()
	err := service.stockRepo.GetStockByProductAndCountry(stock)

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}

	// if no stock found, create new one
	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {

		stock.Count = request.Quantity
		if stock.Count < 0 {
			stock.Count = 0
		}
		return service.stockRepo.SaveStock(stock)
	}

	// if stock found, update it
	stock.Count = stock.Count + request.Quantity
	if stock.Count < 0 {
		stock.Count = 0
	}
	return service.stockRepo.UpdateStock(stock)
}
