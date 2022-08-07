package country

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type CountryItemDTO struct {
	CountryID   uint
	CountryCode string
}

func (dto *CountryItemDTO) toResponseDTO(country *domain.Country) {
	dto.CountryID = country.ID
	dto.CountryCode = country.Code
}
