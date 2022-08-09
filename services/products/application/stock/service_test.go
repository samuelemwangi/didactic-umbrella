package stock

import (
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/mock_persistence"
)

func TestConsumeStock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := mock_persistence.NewMockRepositories(mockCtrl)

	stockService := &stockService{
		stockRepo:    mockRepos.StockRepo,
		errorService: errorhelper.NewErrorService(),
	}

	// Consume Stock - valid request returns valid stock response
	t.Run("Test ConsumeStock() - valid request returns valid stock response", func(t *testing.T) {

		stock := &domain.Stock{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(stock).Return(nil)
		mockRepos.StockRepo.EXPECT().UpdateStock(stock).Return(nil)

		stockRequest := &ConsumeStockRequestDTO{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}

		stockResponse, errResponse := stockService.ConsumeStock(stockRequest)

		if errResponse != nil {
			t.Errorf("Error consuming stock: %s", errResponse.Message)
		}

		if stockResponse != nil && stockResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, stockResponse.Status)
		}

		if stockResponse != nil && !strings.Contains(strings.ToLower(stockResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", stockResponse.Message)
		}

		if stockResponse != nil && stockResponse.Item.QuantityBalance != stock.Quantity {
			t.Errorf("Expected Quantity balance to be %d, got %d", stock.Quantity, stockResponse.Item.QuantityBalance)
		}
		if stockResponse != nil && stockResponse.Item.ProductID != stock.ProductID {
			t.Errorf("Expected ProductID to be %d, got %d", stock.ProductID, stockResponse.Item.ProductID)
		}
		if stockResponse != nil && stockResponse.Item.CountryID != stock.CountryID {
			t.Errorf("Expected CountryID to be %d, got %d", stock.CountryID, stockResponse.Item.CountryID)
		}
	})

}
