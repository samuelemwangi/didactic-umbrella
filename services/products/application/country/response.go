package country

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type countryDetailDTO struct {
	CountryID   uint   `json:"countryId"`
	CountryCode string `json:"countrName"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type countrriesDTO struct {
	CountryID   uint   `json:"countryId"`
	CountryCode string `json:"countrName"`
}

type CountryResponseDTO struct {
	Status  int               `json:"responseStatus"`
	Message string            `json:"responseMessage"`
	Item    *countryDetailDTO `json:"itemDetails"`
}

type CountriesResponseDTO struct {
	Status  int              `json:"responseStatus"`
	Message string           `json:"responseMessage"`
	Items   []*countrriesDTO `json:"items"`
}

func (response *CountryResponseDTO) toCountryResponseDTO(entity *domain.Country) {

	productDetail := &countryDetailDTO{
		CountryID:   entity.ID,
		CountryCode: entity.Code,
		CreatedAt:   entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = productDetail
}

func (response *CountriesResponseDTO) toCountriesResponseDTO(entities []*domain.Country) {

	products := make([]*countrriesDTO, len(entities))
	for i, entity := range entities {
		product := &countrriesDTO{
			CountryID:   entity.ID,
			CountryCode: entity.Code,
		}
		products[i] = product
	}
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Items = products
}
