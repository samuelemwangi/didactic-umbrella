package product

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type ProductService interface {
	SaveProduct(*ProductRequestDTO) (*ProductItemDTO, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(repos *persistence.Repositories) *productService {
	return &productService{
		productRepo: repos.ProductRepo,
	}
}

func (service *productService) SaveProduct(request *ProductRequestDTO) (*ProductItemDTO, error) {
	var responseDTO ProductItemDTO

	product := request.toEntity()
	err := service.productRepo.GetProduct(product)

	if err != nil {
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			err = service.productRepo.SaveProduct(product)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	responseDTO.toResponseDTO(product)

	return &responseDTO, nil

}
