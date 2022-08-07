package stock

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type StockService interface {
	ConsumeStock(*ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *errorhelper.ErrorResponseDTO)
}

type stockService struct {
	stockRepo    repositories.StockRepository
	errorService errorhelper.ErrorService
}

func NewStockService(repos *persistence.Repositories) *stockService {
	return &stockService{
		stockRepo:    repos.StockRepo,
		errorService: errorhelper.NewErrorService(),
	}
}

func (service *stockService) ConsumeStock(request *ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *errorhelper.ErrorResponseDTO) {

	// validate request
	validationErrors := request.validateRequest()

	if len(validationErrors) > 0 {
		return nil, service.errorService.GetValidationError(http.StatusBadRequest, validationErrors)
	}

	stock := request.toEntity()

	// validate stock item exists
	dbError := service.stockRepo.GetStockByProductAndCountry(stock)

	if dbError != nil {
		status := http.StatusInternalServerError

		if dbError.Error() == gorm.ErrRecordNotFound.Error() {
			status = http.StatusNotFound
		}
		return nil, service.errorService.GetGeneralError(status, dbError)
	}

	// validate there is available quantity to consume
	if stock.Quantity < request.Quantity {
		err := errors.New("Not enough stock available. only " + fmt.Sprint(stock.Quantity) + " items available.")
		return nil, service.errorService.GetGeneralError(http.StatusBadRequest, err)
	}

	// update stock count
	stock.Quantity = stock.Quantity - request.Quantity
	updateError := service.stockRepo.UpdateStock(stock)

	log.Println("stock available: " + fmt.Sprint(stock.Quantity))
	log.Println("stock consumed: " + fmt.Sprint(request.Quantity))

	if updateError != nil {
		return nil, service.errorService.GetGeneralError(http.StatusInternalServerError, updateError)
	}

	// prepare response
	var consumeStockResponse ConsumeStockResponseDTO
	consumeStockResponse.toResponseDTO(stock)

	return &consumeStockResponse, nil

}
