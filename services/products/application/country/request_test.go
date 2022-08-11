package country

import (
	"testing"
)

func TestCountryRequestToEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		countryRequest := CountryRequestDTO{
			CountryCode: "KE",
		}

		country := countryRequest.toEntity()

		if country.Code != "KE" {
			t.Errorf("Expected country code to be KE, got %s", country.Code)
		}
	})
}

func TestCountryRequestValidateRequest(t *testing.T) {
	// Invalid request
	t.Run("Test validateRequest() method - Invalid request", func(t *testing.T) {
		countryRequest := CountryRequestDTO{
			CountryCode: "",
		}

		errors := countryRequest.validateRequest()

		if len(errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(errors))
		}

		if errors["CountryCode"] != "required" {
			t.Errorf("Expected error to be required, got %s", errors["CountryCode"])
		}
	})

	// Valid request
	t.Run("Test validateRequest() method - Valid request", func(t *testing.T) {
		countryRequest := CountryRequestDTO{
			CountryCode: "KE",
		}

		errors := countryRequest.validateRequest()

		if len(errors) != 0 {
			t.Errorf("Expected 0 error, got %d", len(errors))
		}
	})
}
