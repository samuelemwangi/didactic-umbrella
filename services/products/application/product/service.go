package product

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type ProductService interface {
	GetProductBySKU(string) (*ProductResponseDTO, *error.ErrorResponseDTO)
}

type productService struct {
	productRepo  repositories.ProductRepository
	errorService error.ErrorService
}

func NewProductService(productRepo repositories.ProductRepository) *productService {
	return &productService{
		productRepo:  productRepo,
		errorService: error.NewErrorService(),
	}
}

func (ps *productService) GetProductBySKU(sku string) (*ProductResponseDTO, *error.ErrorResponseDTO) {

	data, dbError := ps.productRepo.GetProductBySKU(sku)

	if dbError != nil {

		statusCode := http.StatusInternalServerError
		if strings.Contains(*dbError, gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound
		}

		return nil, ps.errorService.GetGeneralError(statusCode, *dbError)
	}

	var productResponse ProductResponseDTO
	productResponse.toResponseDTO(data)

	return &productResponse, nil

}
