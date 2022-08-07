package product

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductDetailDTO struct {
	ID   uint   `json:"id"`
	SKU  string `json:"sku"`
	Name string `json:"productName"`
}

type ProductResponseDTO struct {
	Status  int               `json:"responseStatus"`
	Message string            `json:"responseMessage"`
	Item    *ProductDetailDTO `json:"itemDetails"`
}

func (response *ProductResponseDTO) toResponseDTO(entity *domain.Product) {

	productDetail := &ProductDetailDTO{
		ID:   entity.ID,
		SKU:  entity.SKU,
		Name: entity.Name,
	}
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = productDetail
}
