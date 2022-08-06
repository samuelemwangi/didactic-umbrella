package product

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type ProductRequestDTO struct {
	Name string
	SKU  string
}

func (request *ProductRequestDTO) toEntity() *domain.Product {
	return &domain.Product{
		Name: request.Name,
		SKU:  request.SKU,
	}
}
