package integrationtests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/presentation"
)

func TestGetProcessingStatus(t *testing.T) {
	db := OpenTestDBConnection()
	defer db.Close()

	repos := persistence.NewRepositories(db)
	services := application.NewServices(repos)
	handlers := presentation.NewHandlers(services)

	insertUploadMetadataData(db)
	defer clearUploadMetadataData(db)

	t.Run("GetProcessingStatus - Successful request", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()

		r.GET("/api/v1/upload-status/:uploadId", handlers.FileProcessingHandler.GetProcessingStatus)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/upload-status/12345687", nil)

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

func insertUploadMetadataData(db *gorm.DB) {
	uploadMetadata := domain.UploadMetadata{
		UploadID:           "12345687",
		FileName:           "test.csv",
		ProcessedStatus:    domain.UploadStatusProcessing,
		TotalItems:         10,
		TotalItemsProcesed: 0,
	}
	db.Create(&uploadMetadata)
}

func clearUploadMetadataData(db *gorm.DB) {
	db.Unscoped().Where("upload_id = ?", "12345687").Delete(&domain.UploadMetadata{})
}
