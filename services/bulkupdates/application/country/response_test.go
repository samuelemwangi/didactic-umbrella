package country

import (
	"testing"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

func TestCountryResponseToDTO(t *testing.T) {
	t.Run("Test toDTO() method", func(t *testing.T) {
		country := &domain.Country{
			Code: "KE",
		}
		country.ID = 1

		countryResponse := CountryItemDTO{}
		countryResponse.toResponseDTO(country)

		if countryResponse.CountryID != country.ID {
			t.Errorf("Expected CountryID to be %d, got %d", country.ID, countryResponse.CountryID)
		}

		if countryResponse.CountryCode != country.Code {
			t.Errorf("Expected Code to be %s, got %s", country.Code, countryResponse.CountryCode)
		}
	})

}
