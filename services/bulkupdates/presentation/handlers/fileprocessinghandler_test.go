package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/uploadprocess"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/country_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/product_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/stock_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/uploadprocess_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/infrastructure_mock/queueing_mock"
)

func TestGetProcessingStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUploadProcessorService := uploadprocess_mock.NewMockUploadProcessorService(mockCtrl)

	fileProcessingHandler := &FileProcessingHandler{
		CountryService:         country_mock.NewMockCountryService(mockCtrl),
		ProductService:         product_mock.NewMockProductService(mockCtrl),
		StockService:           stock_mock.NewMockStockService(mockCtrl),
		UploadProcessorService: mockUploadProcessorService,
		kafkaConsumer:          queueing_mock.NewMockKafkaConsumer(mockCtrl),
	}

	uploadID := "upload-id"

	t.Run("Test GetProcessingStatus Handler - Successful Request", func(t *testing.T) {

		// mock the response from the service
		uploadRequest := uploadprocess.UploadProcessRequestDTO{
			UploadID: uploadID,
		}

		uploadResponse := uploadprocess.UploadProcessResponseDTO{
			Status:  http.StatusOK,
			Message: "request successful",
		}
		mockUploadProcessorService.EXPECT().GetProcessingStatus(&uploadRequest).Return(&uploadResponse, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/upload-status/"+uploadID, nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/upload-status/:uploadId", fileProcessingHandler.GetProcessingStatus)
		r.ServeHTTP(rr, req)

		actualResponse := uploadprocess.UploadProcessResponseDTO{}

		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)
		if err != nil {
			t.Errorf("Expected error to be nil, got: %s", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code to be %d, got: %d", http.StatusOK, rr.Code)
		}

		if actualResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be %d, got: %d", http.StatusOK, actualResponse.Status)
		}
	})

	t.Run("Test GetProcessingStatus Handler - Service error returns an error response", func(t *testing.T) {

		// mock the response from the service
		uploadRequest := uploadprocess.UploadProcessRequestDTO{
			UploadID: uploadID,
		}

		errorResponse := errorhelper.ErrorResponseDTO{
			Status:  http.StatusInternalServerError,
			Message: "request failed",
		}

		mockUploadProcessorService.EXPECT().GetProcessingStatus(&uploadRequest).Return(nil, &errorResponse)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/upload-status/"+uploadID, nil)
		rr := httptest.NewRecorder()
		r.GET("/api/v1/upload-status/:uploadId", fileProcessingHandler.GetProcessingStatus)
		r.ServeHTTP(rr, req)

		actualResponse := uploadprocess.UploadProcessResponseDTO{}

		err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)

		if err != nil {
			t.Errorf("Expected error to be nil, got: %s", err)
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code to be %d, got: %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
