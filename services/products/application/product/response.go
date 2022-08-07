package product

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type productDetailDTO struct {
	ID        uint   `json:"id"`
	SKU       string `json:"sku"`
	Name      string `json:"productName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type productsDTO struct {
	ID   uint   `json:"id"`
	SKU  string `json:"sku"`
	Name string `json:"productName"`
}

type ProductResponseDTO struct {
	Status  int               `json:"responseStatus"`
	Message string            `json:"responseMessage"`
	Item    *productDetailDTO `json:"itemDetails"`
}

type ProductsResponseDTO struct {
	Status  int            `json:"responseStatus"`
	Message string         `json:"responseMessage"`
	Items   []*productsDTO `json:"items"`
}

func (response *ProductResponseDTO) toProductResponseDTO(entity *domain.Product) {

	productDetail := &productDetailDTO{
		ID:        entity.ID,
		SKU:       entity.SKU,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = productDetail
}

func (response *ProductsResponseDTO) toProductsResponseDTO(entities []*domain.Product) {

	products := make([]*productsDTO, len(entities))
	for i, entity := range entities {
		product := &productsDTO{
			ID:   entity.ID,
			SKU:  entity.SKU,
			Name: entity.Name,
		}
		products[i] = product
	}
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Items = products
}
