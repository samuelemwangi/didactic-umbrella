package integrationtests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation"
)

func TestGetProducts(t *testing.T) {
	db := OpenTestDBConnection()
	defer db.Close()

	repos := persistence.NewRepositories(db)
	services := application.NewServices(repos)
	handlers := presentation.NewHandlers(services)
	insertProductsData(db)
	defer clearProductsData(db)

	t.Run("GetProducts Test - Get All Products", func(t *testing.T) {

		gin.SetMode(gin.TestMode)
		r := gin.Default()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/products", handlers.ProductHandler.GetProducts)
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err := json.Unmarshal(rr.Body.Bytes(), &responseMap)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, rr.Code)
		}

		if responseMap["responseStatus"] != float64(http.StatusOK) {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, responseMap["responseStatus"])
		}

		if reflect.ValueOf(responseMap["items"]).Len() != 1 {
			t.Errorf("expected number of items %v, actual number of items %v", 1, reflect.ValueOf(responseMap["items"]).Len())
		}

	})

	t.Run("GetProducts Test - Get Product by SKU", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()

		r.GET("/api/v1/products/:sku", handlers.ProductHandler.GetProductBySKU)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/1sku23", nil)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err := json.Unmarshal(rr.Body.Bytes(), &responseMap)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, rr.Code)
		}

		if responseMap["responseStatus"] != float64(http.StatusOK) {
			t.Errorf("expected status code %v, actual status code %v", http.StatusOK, responseMap["responseStatus"])
		}

		if reflect.ValueOf(responseMap["itemDetails"]).Len() != 5 {
			t.Errorf("expected number of items %v, actual number of items %v", 1, reflect.ValueOf(responseMap["itemDetails"]).Len())
		}
	})

}

func insertProductsData(db *gorm.DB) {
	products := domain.Product{
		Name: "Sample Product",
		SKU:  "1sku23",
	}

	db.Create(&products)
}

func clearProductsData(db *gorm.DB) {
	db.Unscoped().Where("sku = ?", "1sku23").Delete(&domain.Product{})
}
