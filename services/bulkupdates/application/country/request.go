package country

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type CountryRequestDTO struct {
	CountryName string
}

func (request *CountryRequestDTO) toEntity() *domain.Country {
	return &domain.Country{
		Name: request.CountryName,
	}
}
