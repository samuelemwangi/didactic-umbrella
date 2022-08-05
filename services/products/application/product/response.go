package product

import "github.com/samuelemwangi/jumia-mds-test/services/products/domain"

type ProductResponseDTO struct {
	ID   uint64 `json:"id"`
	SKU  string `json:"sku"`
	Name string `json:"productName"`
}

func (p *ProductResponseDTO) toResponseDTO(entity *domain.Product) {
	p.ID = entity.ID
	p.SKU = entity.SKU
	p.Name = entity.Name
}
