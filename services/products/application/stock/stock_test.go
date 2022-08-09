package stock

import (
	"net/http"
	"strings"
	"testing"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

// ============================= Test request.go

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
	},
	)
}

func TestInvalidConsumeStockRequestDTOValidateRequest(t *testing.T) {
	t.Run("Test validateRequest() method with an invalid request", func(t *testing.T) {
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
	},
	)
}

func TestValidConsumeStockRequestDTOValidateRequest(t *testing.T) {
	t.Run("Test validateRequest() method with a valid request", func(t *testing.T) {
		request := ConsumeStockRequestDTO{
			ProductID: 1,
			Quantity:  1,
			CountryID: 1,
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
func TestStockResponseToDTO(t *testing.T) {
	t.Run("Test toStockResponseDTO() method", func(t *testing.T) {

		stock := &domain.Stock{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}

		consumeStockResponse := ConsumeStockResponseDTO{}
		consumeStockResponse.toConsumeStockResponseDTO(stock)

		if consumeStockResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", consumeStockResponse.Status)
		}

		if consumeStockResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", consumeStockResponse.Message)
		}

		if consumeStockResponse.Item.QuantityBalance != stock.Quantity {
			t.Errorf("Expected Quantity balance to be %d, got %d", stock.Quantity, consumeStockResponse.Item.QuantityBalance)
		}
		if consumeStockResponse.Item.ProductID != stock.ProductID {
			t.Errorf("Expected ProductID to be %d, got %d", stock.ProductID, consumeStockResponse.Item.ProductID)
		}
		if consumeStockResponse.Item.CountryID != stock.CountryID {
			t.Errorf("Expected CountryID to be %d, got %d", stock.CountryID, consumeStockResponse.Item.CountryID)
		}
	},
	)
}
