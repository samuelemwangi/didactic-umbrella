package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/application_mock/country_mock"
)

func TestGetCountries(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCountryService := country_mock.NewMockCountryService(mockCtrl)

	countryHander := &CountryHandler{
		countryService: mockCountryService,
	}

	t.Run("Get Countries Handler - valid request has valid response", func(t *testing.T) {
		// mock get countries
		countriesresponse := &country.CountriesResponseDTO{
			Status:  http.StatusOK,
			Message: "request successful",
		}

		mockCountryService.EXPECT().GetCountries().Return(countriesresponse, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/countries", nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/countries", countryHander.GetCountries)
		r.ServeHTTP(rr, req)

		actualResponse := country.CountriesResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, rr.Code)
		}

		if actualResponse.Status != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, actualResponse.Status)
		}
	})

	t.Run("Get Countries Handler - service error returns an error response", func(t *testing.T) {
		// mock get countries
		errorResponse := &errorhelper.ErrorResponseDTO{
			Status:  http.StatusInternalServerError,
			Message: "request failed",
		}

		mockCountryService.EXPECT().GetCountries().Return(nil, errorResponse)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/countries", nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/countries", countryHander.GetCountries)
		r.ServeHTTP(rr, req)

		actualResponse := errorhelper.ErrorResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, rr.Code)
		}

		if actualResponse.Status != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, actualResponse.Status)
		}

	})

}

func TestSaveCountry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCountryService := country_mock.NewMockCountryService(mockCtrl)

	countryHander := &CountryHandler{
		countryService: mockCountryService,
	}

	t.Run("Save Country Handler - valid request has valid response", func(t *testing.T) {
		countryCode := "KE"
		// mock save country
		countryRequest := &country.CountryRequestDTO{
			CountryCode: countryCode,
		}

		countryresponse := &country.CountryResponseDTO{
			Status:  http.StatusOK,
			Message: "request successful",
		}

		mockCountryService.EXPECT().SaveCountry(countryRequest).Return(countryresponse, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()

		jsonBody := []byte(`{"countryCode": "` + countryCode + `"}`)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/countries", bytes.NewBuffer(jsonBody))
		rr := httptest.NewRecorder()

		r.POST("/api/v1/countries", countryHander.SaveCountry)
		r.ServeHTTP(rr, req)

		actualResponse := country.CountryResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %v, actual status code %v", http.StatusCreated, rr.Code)
		}

		if actualResponse.Status != http.StatusCreated {
			t.Errorf("expected status code %v, actual status code %v", http.StatusCreated, actualResponse.Status)
		}
	})

	t.Run("Save Country Handler - service error returns an error response", func(t *testing.T) {
		countryCode := "KE"
		// mock save country
		countryRequest := &country.CountryRequestDTO{
			CountryCode: countryCode,
		}

		errorResponse := &errorhelper.ErrorResponseDTO{
			Status:  http.StatusInternalServerError,
			Message: "request failed",
		}

		mockCountryService.EXPECT().SaveCountry(countryRequest).Return(nil, errorResponse)

		gin.SetMode(gin.TestMode)
		r := gin.Default()

		jsonBody := []byte(`{"countryCode": "` + countryCode + `"}`)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/countries", bytes.NewBuffer(jsonBody))
		rr := httptest.NewRecorder()

		r.POST("/api/v1/countries", countryHander.SaveCountry)
		r.ServeHTTP(rr, req)

		actualResponse := errorhelper.ErrorResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, rr.Code)
		}

		if actualResponse.Status != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, actualResponse.Status)
		}
	})
}
