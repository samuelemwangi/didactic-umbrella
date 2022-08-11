package uploadprocess

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/country_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/product_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/stock_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/infrastructure_mock/fileutils_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock"
)

func GetNewTestUploadProcessorService(t *testing.T) (*uploadProcessorService, *persistence_mock.MockRepositories) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)

	return &uploadProcessorService{
		uploadMetdataRepo: mockRepos.UploadMetadataRepo,
		csvReader:         fileutils_mock.NewMockCSVReader(mockCtrl),
		countryService:    country_mock.NewMockCountryService(mockCtrl),
		productService:    product_mock.NewMockProductService(mockCtrl),
		stockService:      stock_mock.NewMockStockService(mockCtrl),
		errorService:      errorhelper.NewErrorService(),
		uploadMetadata:    &domain.UploadMetadata{},
	}, mockRepos
}

func TestGetProcessingStatus(t *testing.T) {

	uploadProcessorService, mockRepos := GetNewTestUploadProcessorService(t)

	uploadId := "sample1234"
	fileName := "sample.csv"
	totalItems := uint(10)
	totalItemsProcessed := uint(5)

	t.Run("Test GetProcessingStatus() method - valid request returns a valid response", func(t *testing.T) {

		uploadMetadata := &domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         totalItems,
			TotalItemsProcesed: totalItemsProcessed,
		}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, nil)

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}

		if uploadResponse != nil && uploadResponse.Status != http.StatusOK {
			t.Errorf("Expected status code %v, but got %v", http.StatusOK, uploadResponse.Status)
		}

		if uploadResponse != nil && uploadResponse.Item.UploadID != uploadId {
			t.Errorf("Expected upload id %v, but got %v", uploadId, uploadResponse.Item.UploadID)
		}
	})

	t.Run("Test GetProcessingStatus() method - invalid request returns an error response", func(t *testing.T) {

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: "",
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if uploadResponse != nil && uploadResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status code %v, but got %v", http.StatusBadRequest, uploadResponse.Status)
		}
	})

	t.Run("Test GetProcessingStatus() method - record not found error returns an error response", func(t *testing.T) {
		uploadMetadata := &domain.UploadMetadata{}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, gorm.ErrRecordNotFound)

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if uploadResponse != nil {
			t.Errorf("Expected nil response, but got %v", uploadResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusNotFound {
			t.Errorf("Expected status code %v, but got %v", http.StatusNotFound, errResponse.Status)
		}
	})

	t.Run("Test GetProcessingStatus() method - db error returns an error response", func(t *testing.T) {
		uploadMetadata := &domain.UploadMetadata{}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, errors.New("db error"))

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if uploadResponse != nil {
			t.Errorf("Expected nil response, but got %v", uploadResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, errResponse.Status)
		}
	})

}

func TestProcessUpload(t *testing.T) {
	// uploadProcessorService, mockRepos := GetNewTestUploadProcessorService(t)

	// uploadId := "sample1234"
	// fileName := "sample.csv"

	t.Run("Test ProcessUpload() method - valid request returns a valid response", func(t *testing.T) {

	})

}
