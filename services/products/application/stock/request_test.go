package stock

import (
	"strings"
	"testing"
)

func TestConsumeStockRequestDTOtoEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		request := ConsumeStockRequestDTO{
			ProductID: 1,
			Quantity:  1,
			CountryID: 1,
		}
		stock := request.toEntity()

		if stock.ProductID != 1 {
			t.Errorf("Expected ProductID to be 1, got %d", stock.ProductID)
		}

		// we do not set to a value greater than 0 since we use it to search for stock item
		if stock.Quantity != 0 {
			t.Errorf("Expected Quantity to be 0, got %d", stock.Quantity)
		}
		if stock.CountryID != 1 {
			t.Errorf("Expected CountryID to be 1, got %d", stock.CountryID)
		}
	})
}

func TestConsumeStockRequestDTOValidateRequest(t *testing.T) {
	// Invalid request
	t.Run("Test validateRequest() method - Invalid request", func(t *testing.T) {
		request := ConsumeStockRequestDTO{
			Quantity: -90,
		}
		errors := request.validateRequest()

		if len(errors) != 3 {
			t.Errorf("Expected 3 errors, got %d", len(errors))
		}

		if errors["ProductID"] != "required" {
			t.Errorf("Expected error message to be required, got %s", errors["ProductID"])
		}
		if !strings.Contains(errors["Quantity"], "greater than or equal") {
			t.Errorf("Expected error message to contain 'greater than or equal', got %s", errors["Quantity"])
		}
		if errors["CountryID"] != "required" {
			t.Errorf("Expected error message to be required, got %s", errors["CountryID"])
		}
	})

	// Valid request
	t.Run("Test validateRequest() method - Valid request", func(t *testing.T) {
		request := ConsumeStockRequestDTO{
			ProductID: 1,
			Quantity:  1,
			CountryID: 1,
		}
		errors := request.validateRequest()

		if len(errors) != 0 {
			t.Errorf("Expected 0 error, got %d", len(errors))
		}
	})
}
