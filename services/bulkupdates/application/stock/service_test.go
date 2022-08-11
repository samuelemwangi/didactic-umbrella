package stock

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock"
)

func TestSaveStock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)

	stockService := &stockService{
		stockRepo: mockRepos.StockRepo,
	}

	t.Run("Test SaveStock() - Updating existing stock returns a valid response", func(t *testing.T) {

		var countryID uint = 1
		var productID uint = 2
		originalQuantity := 20

		requestQuantity := 10

		stock := &domain.Stock{
			CountryID: countryID,
			ProductID: productID,
			Quantity:  originalQuantity,
		}
		stock.ID = 1
		stock.CreatedAt = time.Now()
		stock.UpdatedAt = time.Now()

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(countryID, productID).Return(stock, nil)

		mockRepos.StockRepo.EXPECT().UpdateStock(stock).Return(nil)

		errResponse := stockService.SaveStock(countryID, productID, requestQuantity)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}

	})

	t.Run("Test SaveStock() - Saving a non-existent stock returns a valid response", func(t *testing.T) {
		var countryID uint = 1
		var productID uint = 2

		requestQuantity := 10

		stock := &domain.Stock{}

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(countryID, productID).Return(stock, gorm.ErrRecordNotFound)

		stock.CountryID = countryID
		stock.ProductID = productID
		stock.Quantity = requestQuantity

		mockRepos.StockRepo.EXPECT().SaveStock(stock).Return(nil)

		errResponse := stockService.SaveStock(countryID, productID, requestQuantity)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}
	})

	t.Run("Test SaveStock() - Db error when getting stock returns an error response", func(t *testing.T) {
		var countryID uint = 1
		var productID uint = 2

		requestQuantity := 10

		stock := &domain.Stock{}

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(countryID, productID).Return(stock, errors.New("Db error"))

		errResponse := stockService.SaveStock(countryID, productID, requestQuantity)

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

	})

	t.Run("Test SaveStock() - Db error when updating stock returns an error response", func(t *testing.T) {

		var countryID uint = 1
		var productID uint = 2
		originalQuantity := 20

		requestQuantity := 10

		stock := &domain.Stock{
			CountryID: countryID,
			ProductID: productID,
			Quantity:  originalQuantity,
		}
		stock.ID = 1
		stock.CreatedAt = time.Now()
		stock.UpdatedAt = time.Now()

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(countryID, productID).Return(stock, nil)

		mockRepos.StockRepo.EXPECT().UpdateStock(stock).Return(errors.New("Db error"))

		errResponse := stockService.SaveStock(countryID, productID, requestQuantity)

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

	})

	t.Run("Test SaveStock() - Db error when saving stock returns an error response", func(t *testing.T) {

		var countryID uint = 1
		var productID uint = 2

		requestQuantity := 10

		stock := &domain.Stock{}

		mockRepos.StockRepo.EXPECT().GetStockByProductAndCountry(countryID, productID).Return(stock, gorm.ErrRecordNotFound)

		stock.CountryID = countryID
		stock.ProductID = productID
		stock.Quantity = requestQuantity

		mockRepos.StockRepo.EXPECT().SaveStock(stock).Return(errors.New("Db error"))

		errResponse := stockService.SaveStock(countryID, productID, requestQuantity)

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

	})
}
