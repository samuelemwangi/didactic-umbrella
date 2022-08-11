package product

import (
	"testing"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

func TestProductResponseToDTO(t *testing.T) {

	t.Run("Test toDTO() method", func(t *testing.T) {
		product := &domain.Product{
			Name: "product",
		}
		product.ID = 1
		product.SKU = "sku123"

		productResponse := ProductItemDTO{}
		productResponse.toResponseDTO(product)

		if productResponse.ProductID != product.ID {
			t.Errorf("Expected ProductID to be %d, got %d", product.ID, productResponse.ProductID)
		}

		if productResponse.ProductName != product.Name {
			t.Errorf("Expected Name to be %s, got %s", product.Name, productResponse.ProductName)
		}

		if productResponse.SKU != product.SKU {
			t.Errorf("Expected SKU to be %s, got %s", product.SKU, productResponse.SKU)
		}
	})

}
