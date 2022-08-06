package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type StockService interface {
	ConsumeStock(*ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *error.ErrorResponseDTO)
}

type stockService struct {
	stockRepo    repositories.StockRepository
	errorService error.ErrorService
}

func NewStockService(repos *persistence.Repositories) *stockService {
	return &stockService{
		stockRepo:    repos.StockRepo,
		errorService: error.NewErrorService(),
	}
}

func (sc *stockService) ConsumeStock(request *ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *error.ErrorResponseDTO) {

	// validate request
	validationErrors := request.validateRequest()
	if len(validationErrors) > 0 {
		return nil, sc.errorService.GetValidationError(http.StatusBadRequest, "validation errors occured", validationErrors)
	}

	// validate count in the db
	data, dbError := sc.stockRepo.GetStockByProductAndCountry(request.toEntity())
	if dbError != nil {
		return nil, error.NewErrorService().GetGeneralError(http.StatusBadRequest, *dbError)
	}
	if data.Count < request.Count {
		return nil, error.NewErrorService().GetGeneralError(http.StatusBadRequest, "stock not available")
	}

	// update stock count
	newStock := data.Count - request.Count

	data.Count = newStock

	_, updateError := sc.stockRepo.UpdateStockCount(data)

	if updateError != nil {
		return nil, error.NewErrorService().GetGeneralError(http.StatusInternalServerError, *updateError)
	}

	// prepare response
	var consumeStockResponse ConsumeStockResponseDTO
	consumeStockResponse.toResponseDTO(data)
	return &consumeStockResponse, nil

}
