package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type StockService interface {
	ConsumeStock(*ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *error.ErrorResponseDTO)
}

type stockService struct {
	stockRepo    repositories.StockRepository
	errorService error.ErrorService
}

func NewStockService(stockRepo repositories.StockRepository) *stockService {
	return &stockService{
		stockRepo:    stockRepo,
		errorService: error.NewErrorService(),
	}
}

func (sc *stockService) ConsumeStock(request *ConsumeStockRequestDTO) (*ConsumeStockResponseDTO, *error.ErrorResponseDTO) {
	data, dbError := sc.stockRepo.GetStockByProductAndCountry(request.toEntity())
	if dbError != nil {
		return nil, error.NewErrorService().GetGeneralError(http.StatusBadRequest, *dbError)
	}

	if data.Count < request.Count {
		return nil, error.NewErrorService().GetGeneralError(http.StatusBadRequest, "stock not available")
	}

	newStock := data.Count - request.Count

	data.Count = newStock

	_, updateError := sc.stockRepo.UpdateStockCount(data)

	if updateError != nil {
		return nil, error.NewErrorService().GetGeneralError(http.StatusInternalServerError, *updateError)
	}

	var consumeStockResponse ConsumeStockResponseDTO
	consumeStockResponse.toResponseDTO(data)
	return &consumeStockResponse, nil

}
