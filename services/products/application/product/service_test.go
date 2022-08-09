package product

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/mock_persistence"
)

func TestGetProductBySKU(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := mock_persistence.NewMockRepositories(mockCtrl)

	productService := &productService{
		productRepo:  mockRepos.ProductRepo,
		errorService: errorhelper.NewErrorService(),
	}

	// Get Product By SKU
	t.Run("Test GetProductBySKU() - valid request returns valid product response", func(t *testing.T) {

		product := &domain.Product{
			SKU: "123H",
		}

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(product).Return(nil)

		productRequest := &ProductRequestDTO{
			SKU: "123H",
		}

		productResponse, errResponse := productService.GetProductBySKU(productRequest)

		if errResponse != nil {
			t.Errorf("Error getting product: %s", errResponse.Message)
		}

		if productResponse != nil && productResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, productResponse.Status)
		}

		if productResponse != nil && !strings.Contains(strings.ToLower(productResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", productResponse.Message)
		}

		if productResponse != nil && productResponse.Item.SKU != productRequest.SKU {
			t.Errorf("Expected product to be populated, got %v", productResponse)
		}

	})

	// Invalid request
	t.Run("Test GetProductBySKU() - invalid request returns error response", func(t *testing.T) {

		productRequest := &ProductRequestDTO{
			SKU: "",
		}

		productResponse, errResponse := productService.GetProductBySKU(productRequest)

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if productResponse != nil {
			t.Errorf("Expected product to be nil, got %v", productResponse)
		}

		if errResponse != nil && errResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, errResponse.Status)
		}

	})

	t.Run("Test GetProductBySKU() -  db error  returns error response", func(t *testing.T) {

		product := &domain.Product{
			SKU: "123H",
		}

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(product).Return(errors.New("db error"))

		productRequest := &ProductRequestDTO{
			SKU: "123H",
		}

		productResponse, errResponse := productService.GetProductBySKU(productRequest)

		if productResponse != nil {
			t.Errorf("Expected product to be nil, got %v", productResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, errResponse.Status)
		}

	})

	t.Run("Test GetProductBySKU() -  db error record not found returns a not found error response", func(t *testing.T) {

		product := &domain.Product{
			SKU: "123H",
		}

		mockRepos.ProductRepo.EXPECT().GetProductBySKU(product).Return(gorm.ErrRecordNotFound)

		productRequest := &ProductRequestDTO{
			SKU: "123H",
		}

		productResponse, errResponse := productService.GetProductBySKU(productRequest)

		if productResponse != nil {
			t.Errorf("Expected product to be nil, got %v", productResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, errResponse.Status)
		}

	})
}

func TestGetProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := mock_persistence.NewMockRepositories(mockCtrl)

	productService := &productService{
		productRepo:  mockRepos.ProductRepo,
		errorService: errorhelper.NewErrorService(),
	}

	t.Run("Test GetProducts() - products > 0 returns all products", func(t *testing.T) {
		products := []*domain.Product{
			{
				SKU:  "123H",
				Name: "Product 1",
			},
			{
				SKU:  "456H",
				Name: "Product 2",
			},
		}

		mockRepos.ProductRepo.EXPECT().GetProducts().Return(products, nil)
		countriesResponse, errResponse := productService.GetProducts()
		if errResponse != nil {
			t.Errorf("Error getting products: %s", errResponse.Message)
		}

		if countriesResponse != nil && countriesResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, countriesResponse.Status)
		}

		if countriesResponse != nil && !strings.Contains(strings.ToLower(countriesResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", countriesResponse.Message)
		}

		if countriesResponse != nil && len(countriesResponse.Items) != len(products) {
			t.Errorf("Expected products to be populated, got %v", countriesResponse)
		}
	})

	t.Run("Test GetProducts() - products count = 0 returns empty response", func(t *testing.T) {
		products := []*domain.Product{}

		mockRepos.ProductRepo.EXPECT().GetProducts().Return(products, nil)

		countriesResponse, errResponse := productService.GetProducts()

		if errResponse != nil {
			t.Errorf("Error getting products: %s", errResponse.Message)
		}

		if countriesResponse != nil && countriesResponse.Status != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, countriesResponse.Status)
		}

		if countriesResponse != nil && !strings.Contains(strings.ToLower(countriesResponse.Message), "successful") {
			t.Errorf("Expected message to contain 'successful', got %s", countriesResponse.Message)
		}

		if countriesResponse != nil && len(countriesResponse.Items) != 0 {
			t.Errorf("Expected products to be populated, got %v", countriesResponse)
		}
	})

	t.Run("Test GetProducts() - error getting products returns error response", func(t *testing.T) {
		mockRepos.ProductRepo.EXPECT().GetProducts().Return(nil, errors.New("error getting products"))

		countriesResponse, errResponse := productService.GetProducts()

		if countriesResponse != nil {
			t.Errorf("Expected countries response to be nil, got %v", countriesResponse)
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, errResponse.Status)
		}

		if errResponse != nil && !strings.Contains(strings.ToLower(errResponse.Message), "failed") {
			t.Errorf("Expected message to contain 'failed', got %s", errResponse.Message)
		}
	})
}
