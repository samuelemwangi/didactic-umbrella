package country

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock"
	"gorm.io/gorm"
)

func TestCountryServiceSaveCountry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)

	countryService := &countryService{
		countryRepo: mockRepos.CountryRepo,
	}

	t.Run("Test SaveCountry() method - Existing country record returns valid response", func(t *testing.T) {
		countryCode := "KE"

		country := &domain.Country{
			Code: countryCode,
		}
		country.ID = 1
		country.CreatedAt = time.Now()
		country.UpdatedAt = time.Now()

		mockRepos.CountryRepo.EXPECT().GetCountryByCode(countryCode).Return(country, nil)

		countryResponse, errResponse := countryService.SaveCountry(countryCode)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}

		if countryResponse != nil && countryResponse.CountryCode != country.Code {
			t.Errorf("Expected country code %v, but got %v", country.Code, countryResponse.CountryCode)
		}

		if countryResponse != nil && countryResponse.CountryID != country.ID {
			t.Errorf("Expected country id %v, but got %v", country.ID, countryResponse.CountryID)
		}
	})

	t.Run("Test SaveCountry() method -  Non-existent record returns a valid response", func(t *testing.T) {
		countryCode := "KE"

		country := &domain.Country{}

		mockRepos.CountryRepo.EXPECT().GetCountryByCode(countryCode).Return(country, gorm.ErrRecordNotFound)
		country.Code = countryCode
		mockRepos.CountryRepo.EXPECT().SaveCountry(country).Return(nil)

		countryResponse, errResponse := countryService.SaveCountry(countryCode)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}

		if countryResponse != nil && countryResponse.CountryCode != country.Code {
			t.Errorf("Expected country code %v, but got %v", country.Code, countryResponse.CountryCode)
		}
	})

	t.Run("Test SaveCountry() method -  DB Error when getting record returns an invalid response", func(t *testing.T) {
		countryCode := "KE"
		mockRepos.CountryRepo.EXPECT().GetCountryByCode(countryCode).Return(nil, errors.New("db error"))
		countryResponse, errResponse := countryService.SaveCountry(countryCode)

		if countryResponse != nil {
			t.Errorf("Expected nil response, but got %v", countryResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test SaveCountry() method -  DB Error when saving record returns an invalid response", func(t *testing.T) {
		countryCode := "KE"

		country := &domain.Country{}

		mockRepos.CountryRepo.EXPECT().GetCountryByCode(countryCode).Return(country, gorm.ErrRecordNotFound)
		country.Code = countryCode
		mockRepos.CountryRepo.EXPECT().SaveCountry(country).Return(errors.New("db error"))

		countryResponse, errResponse := countryService.SaveCountry(countryCode)

		if countryResponse != nil {
			t.Errorf("Expected nil response, but got %v", countryResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
