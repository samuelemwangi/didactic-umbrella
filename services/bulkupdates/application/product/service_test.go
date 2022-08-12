package product

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock"
	"gorm.io/gorm"
)

func TestSaveProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)

	productService := &productService{
		productRepo: mockRepos.ProductRepo,
	}

	t.Run("Test SaveProduct() method - Existing product record returns valid response", func(t *testing.T) {
		productSKU := "123"
		productName := "Sample Product"

		product := &domain.Product{
			SKU:  productSKU,
			Name: productName,
		}
		product.ID = 1
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(productSKU).Return(product, nil)

		productResponse, errorResponse := productService.SaveProduct(productSKU, productName)

		if errorResponse != nil {
			t.Errorf("Expected nil error, but got %v", errorResponse)
		}

		if productResponse != nil && productResponse.ProductName != productName {
			t.Errorf("Expected product name %v, but got %v", productName, productResponse.ProductName)
		}

		if productResponse != nil && productResponse.SKU != productSKU {
			t.Errorf("Expected product sku %v, but got %v", productSKU, productResponse.SKU)
		}

	})

	t.Run("Test SaveProduct() method -  Non-existent record returns a valid response", func(t *testing.T) {
		productName := "Sample Product"
		productSKU := "123"

		product := &domain.Product{}

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(productSKU).Return(product, gorm.ErrRecordNotFound)
		product.SKU = productSKU
		product.Name = productName
		mockRepos.ProductRepo.EXPECT().SaveProduct(product).Return(nil)

		productResponse, errorResponse := productService.SaveProduct(productSKU, productName)

		if errorResponse != nil {
			t.Errorf("Expected nil error, but got %v", errorResponse)
		}

		if productResponse != nil && productResponse.ProductName != productName {
			t.Errorf("Expected product name %v, but got %v", productName, productResponse.ProductName)
		}

		if productResponse != nil && productResponse.SKU != productSKU {
			t.Errorf("Expected product sku %v, but got %v", productSKU, productResponse.SKU)
		}

	})

	t.Run("Test SaveProduct() method -  DB Error when getting record returns an error response", func(t *testing.T) {

		productSKU := "123"
		productName := "Sample Product"

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(productSKU).Return(nil, errors.New("db error"))

		productResponse, errorResponse := productService.SaveProduct(productSKU, productName)

		if productResponse != nil {
			t.Errorf("Expected nil product response, but got %v", productResponse)
		}
		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test SaveProduct() method -  DB Error when saving record returns an error response", func(t *testing.T) {

		productSKU := "123"
		productName := "Sample Product"

		product := &domain.Product{}
		mockRepos.ProductRepo.EXPECT().GetProductBySKU(productSKU).Return(product, gorm.ErrRecordNotFound)

		product.SKU = productSKU
		product.Name = productName

		mockRepos.ProductRepo.EXPECT().SaveProduct(product).Return(errors.New("db error"))

		productResponse, errorResponse := productService.SaveProduct(productSKU, productName)

		if productResponse != nil {
			t.Errorf("Expected nil product response, but got %v", productResponse)
		}
		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
