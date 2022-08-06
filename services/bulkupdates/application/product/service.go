package product

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type ProductService interface {
	SaveProduct(*ProductRequestDTO) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(repos *persistence.Repositories) *productService {
	return &productService{
		productRepo: repos.ProductRepo,
	}
}

func (service *productService) SaveProduct(request *ProductRequestDTO) error {
	product := request.toEntity()
	err := service.productRepo.GetProduct(product)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return service.productRepo.SaveProduct(product)
		}
		return err
	}

	return nil

}
