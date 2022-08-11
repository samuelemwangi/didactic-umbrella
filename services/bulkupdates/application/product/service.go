package product

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type ProductService interface {
	SaveProduct(string, string) (*ProductItemDTO, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(repos *persistence.Repositories) *productService {
	return &productService{
		productRepo: repos.ProductRepo,
	}
}

func (service *productService) SaveProduct(sku string, productName string) (*ProductItemDTO, error) {
	var responseDTO ProductItemDTO

	product, err := service.productRepo.GetProductBySKU(sku)

	if err != nil {
		if gorm.ErrRecordNotFound.Error() == err.Error() {

			product.SKU = sku
			product.Name = productName

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
