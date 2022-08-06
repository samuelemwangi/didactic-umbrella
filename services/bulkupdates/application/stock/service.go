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
	err := service.stockRepo.GetStock(stock)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return service.stockRepo.SaveStock(stock)
		}
		return err
	}

	return nil
}
