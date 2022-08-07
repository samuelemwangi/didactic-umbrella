package country

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type CountryService interface {
	SaveCountry(*CountryRequestDTO) (*CountryResponseDTO, *errorhelper.ErrorResponseDTO)
	GetCountries() (*CountriesResponseDTO, *errorhelper.ErrorResponseDTO)
}

type countryService struct {
	countryRepo  repositories.CountryRepository
	errorService errorhelper.ErrorService
}

func NewCountryService(repos *persistence.Repositories) *countryService {
	return &countryService{
		countryRepo:  repos.CountryRepo,
		errorService: errorhelper.NewErrorService(),
	}
}

func (service *countryService) SaveCountry(countryRequest *CountryRequestDTO) (*CountryResponseDTO, *errorhelper.ErrorResponseDTO) {

	// validate request
	validationErrors := countryRequest.validateRequest()

	if len(validationErrors) > 0 {
		return nil, service.errorService.GetValidationError(http.StatusBadRequest, validationErrors)
	}

	// create item
	country := countryRequest.toEntity()
	dbError := service.countryRepo.SaveCountry(country)

	if dbError != nil {
		return nil, service.errorService.GetGeneralError(http.StatusInternalServerError, dbError)

	}
	var countryResponse CountryResponseDTO
	countryResponse.toCountryResponseDTO(country)

	return &countryResponse, nil
}

func (service *countryService) GetCountries() (*CountriesResponseDTO, *errorhelper.ErrorResponseDTO) {

	// get items
	countries, dbError := service.countryRepo.GetCountries()

	if dbError != nil {
		return nil, service.errorService.GetGeneralError(http.StatusInternalServerError, dbError)
	}

	var countriesResponse CountriesResponseDTO
	countriesResponse.toCountriesResponseDTO(countries)

	return &countriesResponse, nil
}
