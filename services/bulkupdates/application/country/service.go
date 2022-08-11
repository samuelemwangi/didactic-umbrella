package country

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type CountryService interface {
	SaveCountry(string) (*CountryItemDTO, error)
}

type countryService struct {
	countryRepo repositories.CountryRepository
}

func NewCountryService(repos *persistence.Repositories) *countryService {
	return &countryService{
		countryRepo: repos.CountryRepo,
	}
}

func (service *countryService) SaveCountry(countryCode string) (*CountryItemDTO, error) {
	var responseDTO CountryItemDTO

	country, err := service.countryRepo.GetCountryByCode(countryCode)

	if err != nil {
		if gorm.ErrRecordNotFound.Error() == err.Error() {

			country.Code = countryCode
			err = service.countryRepo.SaveCountry(country)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	responseDTO.toResponseDTO(country)

	return &responseDTO, nil

}
