package application

import (
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type CountryService interface {
	SaveCountry(*domain.Country) (*domain.Country, map[string]string)
}

type countryService struct {
	cr repositories.CountryRepository
}

func NewCountryService(countryRepo repositories.CountryRepository) *countryService {
	return &countryService{
		cr: countryRepo,
	}
}

func (cs *countryService) SaveCountry(country *domain.Country) (*domain.Country, map[string]string) {
	return cs.cr.SaveCountry(country)
}
