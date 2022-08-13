package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/application_mock/product_mock"
)

func TestGetProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductService := product_mock.NewMockProductService(mockCtrl)

	productHander := &ProductHandler{
		productService: mockProductService,
	}

	t.Run("Get Products Handler - valid request has valid response", func(t *testing.T) {
		// mock get countries
		productsResponse := &product.ProductsResponseDTO{
			Status:  http.StatusOK,
			Message: "request successful",
		}

		mockProductService.EXPECT().GetProducts().Return(productsResponse, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/products", productHander.GetProducts)
		r.ServeHTTP(rr, req)

		actualResponse := product.ProductsResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, rr.Code)
		}

		if actualResponse.Status != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, actualResponse.Status)
		}
	})

	t.Run("Get Products Handler - service error returns an error response", func(t *testing.T) {
		// mock get countries
		errorResponse := &errorhelper.ErrorResponseDTO{
			Status:  http.StatusInternalServerError,
			Message: "request failed",
		}

		mockProductService.EXPECT().GetProducts().Return(nil, errorResponse)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/products", productHander.GetProducts)
		r.ServeHTTP(rr, req)

		actualResponse := errorhelper.ErrorResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, rr.Code)
		}

	})

}

func TestGetProductBySKU(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductService := product_mock.NewMockProductService(mockCtrl)

	productHander := &ProductHandler{
		productService: mockProductService,
	}

	sku := "sku1234"

	t.Run("Get Product By SKU Handler - valid request has valid response", func(t *testing.T) {
		// mock get product by sku
		productRequest := &product.ProductRequestDTO{
			SKU: sku,
		}

		productResponse := &product.ProductResponseDTO{
			Status:  http.StatusOK,
			Message: "request successful",
		}

		mockProductService.EXPECT().GetProductBySKU(productRequest).Return(productResponse, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/"+sku, nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/products/:sku", productHander.GetProductBySKU)
		r.ServeHTTP(rr, req)

		actualResponse := product.ProductResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, rr.Code)
		}

		if actualResponse.Status != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, actualResponse.Status)
		}

	})

	t.Run("Get Product By SKU Handler - service error returns an error response", func(t *testing.T) {
		// mock get product by sku
		productRequest := &product.ProductRequestDTO{
			SKU: sku,
		}

		errorResponse := &errorhelper.ErrorResponseDTO{
			Status:  http.StatusInternalServerError,
			Message: "request failed",
		}

		mockProductService.EXPECT().GetProductBySKU(productRequest).Return(nil, errorResponse)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/"+sku, nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/products/:sku", productHander.GetProductBySKU)
		r.ServeHTTP(rr, req)

		actualResponse := errorhelper.ErrorResponseDTO{}
		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, rr.Code)
		}

		if actualResponse.Status != http.StatusInternalServerError {
			t.Errorf("expected status code %v, actual status code %v", http.StatusInternalServerError, actualResponse.Status)
		}

	})

}
