package product

import (
	"net/http"
	"testing"
	"time"

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

		sku := "123H"
		productName := "Sample Name"

		product := &domain.Product{
			SKU:  sku,
			Name: productName,
		}
		product.ID = 1
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		productResponse := ProductResponseDTO{}
		productResponse.toProductResponseDTO(product)

		if productResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", productResponse.Status)
		}

		if productResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", productResponse.Message)
		}

		if productResponse.Item.ID != product.ID {
			t.Errorf("Expected product id to be %d, got %d", product.ID, productResponse.Item.ID)
		}

		if productResponse.Item.SKU != product.SKU {
			t.Errorf("Expected product sku to be %s, got %s", product.SKU, productResponse.Item.SKU)
		}

		if productResponse.Item.Name != product.Name {
			t.Errorf("Expected product name to be %s, got %s", product.Name, productResponse.Item.Name)
		}

		if productResponse.Item.CreatedAt == (time.Time{}).Format("2006-01-02 15:04:05") {
			t.Errorf("Expected created at to be set, got %s", productResponse.Item.CreatedAt)
		}

		if productResponse.Item.UpdatedAt == (time.Time{}).Format("2006-01-02 15:04:05") {
			t.Errorf("Expected updated at to be set, got %s", productResponse.Item.UpdatedAt)
		}
	},
	)
}

func TestProductsResponseToDTO(t *testing.T) {
	t.Run("Test toProductsResponseDTO() method", func(t *testing.T) {

		sku := "123H"
		productName := "Sample Name"

		product1 := &domain.Product{
			SKU:  sku + "1",
			Name: productName + "1",
		}
		product1.ID = 1

		product2 := &domain.Product{
			SKU:  sku + "2",
			Name: productName + "2",
		}
		product2.ID = 2

		products := []*domain.Product{product1, product2}
		productsResponse := ProductsResponseDTO{}
		productsResponse.toProductsResponseDTO(products)

		if productsResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", productsResponse.Status)
		}

		if productsResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", productsResponse.Message)
		}

		if len(productsResponse.Items) != 2 {
			t.Errorf("Expected 2 products, got %d", len(productsResponse.Items))
		}

		if productsResponse.Items[0].ID != product1.ID {
			t.Errorf("Expected product id to be %d, got %d", product1.ID, productsResponse.Items[0].ID)
		}

		if productsResponse.Items[0].SKU != product1.SKU {
			t.Errorf("Expected product sku to be %s, got %s", product1.SKU, productsResponse.Items[0].SKU)
		}

		if productsResponse.Items[0].Name != product1.Name {
			t.Errorf("Expected product name to be %s, got %s", product1.Name, productsResponse.Items[0].Name)
		}

		if productsResponse.Items[1].ID != product2.ID {
			t.Errorf("Expected product id to be %d, got %d", product2.ID, productsResponse.Items[1].ID)
		}

		if productsResponse.Items[1].SKU != product2.SKU {
			t.Errorf("Expected product sku to be %s, got %s", product2.SKU, productsResponse.Items[1].SKU)
		}

		if productsResponse.Items[1].Name != product2.Name {
			t.Errorf("Expected product name to be %s, got %s", product2.Name, productsResponse.Items[1].Name)
		}
	},
	)
}
