package product

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type ProductItemDTO struct {
	ProductID   uint
	SKU         string
	ProductName string
}

func (dto *ProductItemDTO) toResponseDTO(entity *domain.Product) {
	dto.ProductID = entity.ID
	dto.SKU = entity.SKU
	dto.ProductName = entity.Name
}
