package product

import (
	"testing"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

// ============================= Test request.go

func TestProductRequestDTOToEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		request := ProductRequestDTO{
			SKU: "123H",
		}
		product := request.toEntity()

		if product.SKU != "123H" {
			t.Errorf("Expected SKU to be 123, got %s", product.SKU)
		}
	},
	)
}

func TestInvalidProductRequestDTOValidateRequest(t *testing.T) {
	t.Run("Test validateRequest() method with an invalid request", func(t *testing.T) {
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
	},
	)
}

func TestValidProductRequestDTOValidateRequest(t *testing.T) {
	t.Run("Test validateRequest() method with a valid request", func(t *testing.T) {
		request := ProductRequestDTO{
			SKU: "123H",
		}
		errors := request.validateRequest()

		if len(errors) != 0 {
			t.Errorf("Expected 0 error, got %d", len(errors))
		}
	},
	)
}

// ============================= Test service.go

// ============================= Test response.go
func TestProductResponseToDTO(t *testing.T) {
	t.Run("Test toProductResponseDTO() method", func(t *testing.T) {

		product := &domain.Product{
			SKU: "123H",
		}
		response := ProductResponseDTO{}
		response.toProductResponseDTO(product)

		if response.Item.SKU != "123H" {
			t.Errorf("Expected SKU to be 123, got %s", response.Item.SKU)
		}
	},
	)
}
