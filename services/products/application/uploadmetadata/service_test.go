package uploadmetadata

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/infrastructure_mock/queueing_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/persistence_mock"
)

func TestSaveUploadMetadaData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)
	mockKafkaProducer := queueing_mock.NewMockKafkaProducer(mockCtrl)

	uploadMetadataService := &uploadMetadataService{
		uploadMetadataRepo: mockRepos.UploadMetadataRepo,
		kafkaProducer:      mockKafkaProducer,
		errorService:       errorhelper.NewErrorService(),
		fileIdTopic:        "file-process-topic",
	}

	// Save Upload Metadata Data with valid request
	t.Run("Test SaveUploadMetadaData() - valid request returns valid response", func(t *testing.T) {

		uploadID := uuid.New().String()
		fileName := "test.csv"

		uploadMetadata := &domain.UploadMetadata{
			FileName:        fileName,
			UploadID:        uploadID,
			ProcessedStatus: domain.UploadStatusUploaded,
		}

		mockRepos.UploadMetadataRepo.EXPECT().SaveUploadMetaData(uploadMetadata).Return(nil)
		mockKafkaProducer.EXPECT().ProduceMessage(uploadMetadataService.fileIdTopic, uploadID).Return()

		uploadMetadataRequest := &UploadMetadataDTO{
			FileName: fileName,
			UploadID: uploadID,
		}

		uploadMetadataResponse, errResponse := uploadMetadataService.SaveUploadMetadaData(uploadMetadataRequest)

		if errResponse != nil {
			t.Errorf("Expected nil error response, got %v", errResponse)
		}

		if uploadMetadataResponse.Item.UploadID != uploadID {
			t.Errorf("Expected UploadID %s, got %s", uploadID, uploadMetadataResponse.Item.UploadID)
		}

		if uploadMetadataResponse.Item.UploadedFileName != fileName {
			t.Errorf("Expected FileName %s, got %s", fileName, uploadMetadataResponse.Item.UploadedFileName)
		}

		if uploadMetadataResponse.Item.ProcessingStatus != "New Upload" {
			t.Errorf("Expected ProcessedStatus %s, got %s", "New Upload", uploadMetadataResponse.Item.ProcessingStatus)
		}

		if uploadMetadataResponse.Status != http.StatusOK {
			t.Errorf("Expected Status %d, got %d", http.StatusOK, uploadMetadataResponse.Status)
		}

		if !strings.Contains(uploadMetadataResponse.Message, "successful") {
			t.Errorf("Expected Message %s, got %s", "successful", uploadMetadataResponse.Message)
		}

	})

	// Save Upload Metadata Data has db Error
	t.Run("Test SaveUploadMetadaData() - db error returns error response", func(t *testing.T) {

		uploadID := uuid.New().String()
		fileName := "test.csv"

		uploadMetadata := &domain.UploadMetadata{
			FileName:        fileName,
			UploadID:        uploadID,
			ProcessedStatus: domain.UploadStatusUploaded,
		}

		mockRepos.UploadMetadataRepo.EXPECT().SaveUploadMetaData(uploadMetadata).Return(errors.New("db error"))
		// mockKafkaProducer.EXPECT().ProduceMessage(uploadMetadataService.fileIdTopic, uploadID).Return()

		uploadMetadataRequest := &UploadMetadataDTO{
			FileName: fileName,
			UploadID: uploadID,
		}

		uploadMetadataResponse, errResponse := uploadMetadataService.SaveUploadMetadaData(uploadMetadataRequest)

		if uploadMetadataResponse != nil {
			t.Errorf("Expected nil response, got %v", uploadMetadataResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error response, got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected Status %d, got %d", http.StatusInternalServerError, errResponse.Status)
		}

		if !strings.Contains(errResponse.Message, "failed") {
			t.Errorf("Expected Message %s, got %s", "failed", errResponse.Message)
		}

	})
}
