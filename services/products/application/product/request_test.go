package product

import (
	"testing"
)

func TestProductRequestDTOToEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		request := ProductRequestDTO{
			SKU: "123H",
		}
		product := request.toEntity()

		if product.SKU != "123H" {
			t.Errorf("Expected SKU to be 123, got %s", product.SKU)
		}
	})
}

func TestProductRequestDTOValidateRequest(t *testing.T) {
	// Invalid request
	t.Run("Test validateRequest() method - Invalid request", func(t *testing.T) {
		request := ProductRequestDTO{
			SKU: "",
		}
		errors := request.validateRequest()

		if len(errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(errors))
		}

		if errors["SKU"] != "required" {
			t.Errorf("Expected error message to be required, got %s", errors["SKU"])
		}
	})

	// Valid request
	t.Run("Test validateRequest() method  - Valid request", func(t *testing.T) {
		request := ProductRequestDTO{
			SKU: "123H",
		}
		errors := request.validateRequest()

		if len(errors) != 0 {
			t.Errorf("Expected 0 error, got %d", len(errors))
		}
	})
}
