package stock

import (
	"net/http"
	"testing"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

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
	})
}
