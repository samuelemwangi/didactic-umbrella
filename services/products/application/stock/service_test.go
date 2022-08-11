package stock

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/persistence_mock"
)

func TestConsumeStock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)

	stockService := &stockService{
		stockRepo:    mockRepos.StockRepo,
		errorService: errorhelper.NewErrorService(),
	}

	// Consume Stock - valid request returns valid stock response
	t.Run("Test ConsumeStock() - valid request returns valid stock response", func(t *testing.T) {

		originalStockAmount := 90

		stockRequest := &ConsumeStockRequestDTO{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}

		consumeStockEntity := &domain.Stock{
			CountryID: 1,
			ProductID: 2,
			Quantity:  originalStockAmount,
		}

		consumeStockEntity.ID = 1
		consumeStockEntity.CreatedAt = time.Now()
		consumeStockEntity.UpdatedAt = time.Now()

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(stockRequest.CountryID, stockRequest.ProductID).Return(consumeStockEntity, nil)

		mockRepos.StockRepo.EXPECT().UpdateStock(consumeStockEntity).Return(nil)

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

		if stockResponse != nil && stockResponse.Item.QuantityBalance != consumeStockEntity.Quantity {
			t.Errorf("Expected Quantity balance to be %d, got %d", consumeStockEntity.Quantity, stockResponse.Item.QuantityBalance)
		}
		if stockResponse != nil && stockResponse.Item.ProductID != consumeStockEntity.ProductID {
			t.Errorf("Expected ProductID to be %d, got %d", consumeStockEntity.ProductID, stockResponse.Item.ProductID)
		}
		if stockResponse != nil && stockResponse.Item.CountryID != consumeStockEntity.CountryID {
			t.Errorf("Expected CountryID to be %d, got %d", consumeStockEntity.CountryID, stockResponse.Item.CountryID)
		}

		if stockResponse != nil && stockResponse.Item.QuantityBalance != originalStockAmount-stockRequest.Quantity {
			t.Errorf("Expected Quantity balance to be %d, got %d", originalStockAmount-stockRequest.Quantity, stockResponse.Item.QuantityBalance)
		}
	})

	// Consume Stock - invalid request returns error response
	t.Run("Test ConsumeStock() - invalid request returns error response", func(t *testing.T) {
		stockRequest := &ConsumeStockRequestDTO{}
		stockResponse, errResponse := stockService.ConsumeStock(stockRequest)

		if stockResponse != nil {
			t.Errorf("Expected stock response to be nil, got %v", stockResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, errResponse.Status)
		}

		if errResponse != nil && !strings.Contains(strings.ToLower(errResponse.Message), "validation failed") {
			t.Errorf("Expected message to contain 'validation failed', got %s", errResponse.Message)
		}

	})

	// Consume Stock - more than available returns error response
	t.Run("Test ConsumeStock() - more than available stock returns error response", func(t *testing.T) {
		stockRequest := &ConsumeStockRequestDTO{
			Quantity:  100,
			CountryID: 1,
			ProductID: 2,
		}

		consumeStockEntity := &domain.Stock{
			CountryID: 1,
			ProductID: 2,
			Quantity:  10,
		}

		consumeStockEntity.ID = 1
		consumeStockEntity.CreatedAt = time.Now()
		consumeStockEntity.UpdatedAt = time.Now()

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(stockRequest.CountryID, stockRequest.ProductID).Return(consumeStockEntity, nil)

		stockResponse, errResponse := stockService.ConsumeStock(stockRequest)

		if stockResponse != nil {
			t.Errorf("Expected stock response to be nil, got %v", stockResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, errResponse.Status)
		}

		if errResponse != nil && !strings.Contains(strings.ToLower(errResponse.Message), "request failed") {
			t.Errorf("Expected message to contain 'request failed', got %s", errResponse.Message)
		}

	})

	// Consume Stock - no stock item available returns error response
	t.Run("Test ConsumeStock() - no stock item available returns error response", func(t *testing.T) {
		stockRequest := &ConsumeStockRequestDTO{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}
		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(stockRequest.CountryID, stockRequest.ProductID).Return(nil, gorm.ErrRecordNotFound)
		stockResponse, errResponse := stockService.ConsumeStock(stockRequest)

		if stockResponse != nil {
			t.Errorf("Expected stock response to be nil, got %v", stockResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, errResponse.Status)
		}

		if errResponse != nil && !strings.Contains(strings.ToLower(errResponse.Message), "request failed") {
			t.Errorf("Expected message to contain 'request failed', got %s", errResponse.Message)
		}

	})

	// Consume Stock - db error returns error response
	t.Run("Test ConsumeStock() - db error consuming stock returns error response", func(t *testing.T) {
		stockRequest := &ConsumeStockRequestDTO{
			Quantity:  20,
			CountryID: 1,
			ProductID: 2,
		}
		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(stockRequest.CountryID, stockRequest.ProductID).Return(nil, errors.New("db error"))
		stockResponse, errResponse := stockService.ConsumeStock(stockRequest)

		if stockResponse != nil {
			t.Errorf("Expected stock response to be nil, got %v", stockResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, errResponse.Status)
		}

		if errResponse != nil && !strings.Contains(strings.ToLower(errResponse.Message), "request failed") {
			t.Errorf("Expected message to contain 'request failed', got %s", errResponse.Message)
		}
	})
}
