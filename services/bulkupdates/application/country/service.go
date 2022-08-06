package country

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type CountryService interface {
	SaveCountry(*CountryRequestDTO) error
}

type countryService struct {
	countryRepo repositories.CountryRepository
}

func NewCountryService(repos *persistence.Repositories) *countryService {
	return &countryService{
		countryRepo: repos.CountryRepo,
	}
}

func (service *countryService) SaveCountry(request *CountryRequestDTO) error {
	country := request.toEntity()
	err := service.countryRepo.GetCountry(country)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return service.countryRepo.SaveCountry(country)
		}
		return err
	}

	return nil

}
