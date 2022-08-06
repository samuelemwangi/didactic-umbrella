package country

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type CountryService interface {
	SaveCountry(*CountryRequestDTO) (*CountryResponseDTO, *error.ErrorResponseDTO)
}

type countryService struct {
	countryRepo  repositories.CountryRepository
	errorService error.ErrorService
}

func NewCountryService(repos *persistence.Repositories) *countryService {
	return &countryService{
		countryRepo:  repos.CountryRepo,
		errorService: error.NewErrorService(),
	}
}

func (cs *countryService) SaveCountry(countryRequest *CountryRequestDTO) (*CountryResponseDTO, *error.ErrorResponseDTO) {

	validationErrors := countryRequest.validateRequest()

	if len(validationErrors) > 0 {
		return nil, cs.errorService.GetValidationError(http.StatusBadRequest, "validation errors occured", validationErrors)
	}

	data, dbError := cs.countryRepo.SaveCountry(countryRequest.toEntity())

	if dbError != nil {
		return nil, cs.errorService.GetGeneralError(http.StatusUnprocessableEntity, *dbError)

	}
	countryResponse := CountryResponseDTO{}
	countryResponse.toResponseDTO(data)

	return &countryResponse, nil
}
