package country

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryDetailDTO struct {
	CountryID   uint   `json:"countryId"`
	CountryCode string `json:"countrName"`
}

type CountryResponseDTO struct {
	Status  int               `json:"responseStatus"`
	Message string            `json:"responseMessage"`
	Item    *CountryDetailDTO `json:"itemDetails"`
}

func (c *CountryResponseDTO) toResponseDTO(entity *domain.Country) {

	productDetail := &CountryDetailDTO{
		CountryID:   entity.ID,
		CountryCode: entity.Code,
	}
	c.Status = http.StatusOK
	c.Message = "request successful"
	c.Item = productDetail
}
