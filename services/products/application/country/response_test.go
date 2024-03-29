package country

import (
	"net/http"
	"testing"
	"time"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

func TestCountryResponseToDTO(t *testing.T) {
	t.Run("Test toCountryResponseDTO() method", func(t *testing.T) {

		country := &domain.Country{
			Code: "KE",
		}
		country.ID = 1
		country.CreatedAt = time.Now()
		country.UpdatedAt = time.Now()

		countryResponse := CountryResponseDTO{}
		countryResponse.toCountryResponseDTO(country)

		if countryResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", countryResponse.Status)
		}

		if countryResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", countryResponse.Message)
		}

		if countryResponse.Item.CountryID != 1 {
			t.Errorf("Expected country id to be 1, got %d", countryResponse.Item.CountryID)
		}

		if countryResponse.Item.CountryCode != "KE" {
			t.Errorf("Expected country code to be KE, got %s", countryResponse.Item.CountryCode)
		}

		if countryResponse.Item.CreatedAt == (time.Time{}).Format("2006-01-02 15:04:05") {
			t.Errorf("Expected created at to be set, got %s", countryResponse.Item.CreatedAt)
		}

		if countryResponse.Item.UpdatedAt == (time.Time{}).Format("2006-01-02 15:04:05") {
			t.Errorf("Expected updated at to be set, got %s", countryResponse.Item.UpdatedAt)
		}

	})
}

func TestCountriesResponseToDTO(t *testing.T) {
	t.Run("Test toCountriesResponseDTO() method", func(t *testing.T) {

		countries := []*domain.Country{}
		countries = append(countries, &domain.Country{
			Code: "KE",
		})
		countries = append(countries, &domain.Country{
			Code: "UG",
		})

		countriesResponse := CountriesResponseDTO{}
		countriesResponse.toCountriesResponseDTO(countries)

		if countriesResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", countriesResponse.Status)
		}

		if countriesResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", countriesResponse.Message)
		}

		if len(countriesResponse.Items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(countriesResponse.Items))
		}

		if countriesResponse.Items[0].CountryCode != "KE" {
			t.Errorf("Expected country code to be KE, got %s", countriesResponse.Items[0].CountryCode)
		}

		if countriesResponse.Items[1].CountryCode != "UG" {
			t.Errorf("Expected country code to be UG, got %s", countriesResponse.Items[1].CountryCode)
		}

	})
}
