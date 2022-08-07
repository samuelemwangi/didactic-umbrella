package product

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type ProductItemDTO struct {
	ProductId uint
	SKU       string
	Name      string
}

func (dto *ProductItemDTO) toResponseDTO(entity *domain.Product) {
	dto.ProductId = entity.ID
	dto.SKU = entity.SKU
	dto.Name = entity.Name
}
