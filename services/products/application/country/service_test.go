package country

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/mock_persistence"
)

func TestSaveCountry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := mock_persistence.NewMockRepositories(mockCtrl)

	countryService := &countryService{
		countryRepo:  mockRepos.CountryRepo,
		errorService: errorhelper.NewErrorService(),
	}

	// Save Country
	t.Run("Test SaveCountry() - valid request returns valid country response", func(t *testing.T) {

		country := &domain.Country{
			Code: "KE",
		}

		mockRepos.CountryRepo.EXPECT().SaveCountry(country).Return(nil)

		countryRequest := &CountryRequestDTO{
			CountryCode: "KE",
		}

		countryResponse, errResponse := countryService.SaveCountry(countryRequest)

		if errResponse != nil {
			t.Errorf("Error saving country: %s", errResponse.Message)
		}

		if countryResponse != nil && countryResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, countryResponse.Status)
		}

		if countryResponse != nil && !strings.Contains(strings.ToLower(countryResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", countryResponse.Message)
		}

		if countryResponse != nil && countryResponse.Item.CountryCode != countryRequest.CountryCode {
			t.Errorf("Expected country to be populated, got %v", countryResponse)
		}

	})

	// Invalid request
	t.Run("Test SaveCountry() - invalid request returns error response", func(t *testing.T) {

		countryRequest := &CountryRequestDTO{
			CountryCode: "",
		}

		countryResponse, errResponse := countryService.SaveCountry(countryRequest)

		if countryResponse != nil {
			t.Errorf("Expected country to be nil, got %v", countryResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, errResponse.Status)
		}
	})

}

func TestGetCountries(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := mock_persistence.NewMockRepositories(mockCtrl)
	mockedCountries := []*domain.Country{
		{
			Code: "KE",
		},
		{
			Code: "UG",
		},
	}

	mockedCountries[0].ID = 1
	mockedCountries[0].ID = 1

	countryService := &countryService{
		countryRepo:  mockRepos.CountryRepo,
		errorService: errorhelper.NewErrorService(),
	}

	// Get All Countries
	t.Run("Test GetCountries() - countries > 0 - returns all countries", func(t *testing.T) {

		mockRepos.CountryRepo.EXPECT().GetCountries().Return(mockedCountries, nil)
		countriesResponse, err := countryService.GetCountries()

		if err != nil {
			t.Errorf("Error getting countries: %s", err.Message)
		}

		if len(countriesResponse.Items) != len(mockedCountries) {
			t.Errorf("Expected %d countries, got %d", len(mockedCountries), len(countriesResponse.Items))
		}

		if countriesResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, countriesResponse.Status)
		}

		if !strings.Contains(strings.ToLower(countriesResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", countriesResponse.Message)
		}
	})

	// A database error occurred
	t.Run("Test GetCountries() -  returns error", func(t *testing.T) {

		dbError := errors.New("a database error occured")

		mockRepos.CountryRepo.EXPECT().GetCountries().Return(nil, dbError)

		countriesResponse, errorResponse := countryService.GetCountries()

		if countriesResponse != nil {
			t.Errorf("Expected countries to be nil, got %v", countriesResponse)
		}

		if errorResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errorResponse != nil && errorResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, errorResponse.Status)
		}

		if errorResponse != nil && !strings.Contains(strings.ToLower(errorResponse.Message), "failed") {
			t.Errorf("Expected message to contain 'failed', got %s", errorResponse.Message)
		}
	})

	// No countries found
	t.Run("Test GetCountries() - countries == 0 - returns []", func(t *testing.T) {

		countriesResult := []*domain.Country{}

		mockRepos.CountryRepo.EXPECT().GetCountries().Return(countriesResult, nil)
		countriesResponse, errorResponse := countryService.GetCountries()

		if errorResponse != nil {
			t.Errorf("Expected error response to be nil, got %v", errorResponse)
		}

		if countriesResponse != nil && len(countriesResponse.Items) != 0 {
			t.Errorf("Expected countries count to be 0, got %v", countriesResponse)
		}

		if countriesResponse != nil && countriesResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, countriesResponse.Status)
		}

		if countriesResponse != nil && !strings.Contains(strings.ToLower(countriesResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", countriesResponse.Message)
		}

	})

}
