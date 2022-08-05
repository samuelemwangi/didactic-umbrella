package country

import (
	"fmt"
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductDetail struct {
	ID          string `json:"id"`
	CountryName string `json:"countrName"`
}

type CountryResponseDTO struct {
	Status  int            `json:"responseStatus"`
	Message string         `json:"responseMessage"`
	Item    *ProductDetail `json:"itemDetails"`
}

func (c *CountryResponseDTO) toResponseDTO(entity *domain.Country) {

	productDetail := &ProductDetail{
		ID:          fmt.Sprint(entity.ID),
		CountryName: entity.Name,
	}

	c.Status = http.StatusOK
	c.Message = "request successful"
	c.Item = productDetail
}
