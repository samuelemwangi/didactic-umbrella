package uploadmetadata

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

// ============================= Test request.go
func TestValidUploadMetadataDTOtoEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		request := UploadMetadataDTO{
			FileName: "test.csv",
			UploadID: uuid.New().String(),
		}
		uploadMetadata := request.toEntity()

		if uploadMetadata.FileName != "test.csv" {
			t.Errorf("Expected FileName to be test.csv, got %s", uploadMetadata.FileName)
		}

		if uploadMetadata.UploadID != request.UploadID {
			t.Errorf("Expected UploadID to be %s, got %s", request.UploadID, uploadMetadata.UploadID)
		}

		if uploadMetadata.ProcessedStatus != domain.UploadStatusUploaded {
			t.Errorf("Expected ProcessedStatus to be %d, got %d", domain.UploadStatusUploaded, uploadMetadata.ProcessedStatus)
		}
	},
	)
}

// ============================= Test service.go

// ============================= Test response.go
func TestUploadResponseToDTO(t *testing.T) {
	t.Run("Test toUploadResponseDTO() method", func(t *testing.T) {

		uploadMetdata := &domain.UploadMetadata{
			UploadID:        uuid.New().String(),
			FileName:        "file1.csv",
			ProcessedStatus: domain.UploadStatusUploaded,
		}

		uploadMetdata.CreatedAt = time.Now()

		uploadResponse := UploadResponseDTO{}
		uploadResponse.toUploadResponseDTO(uploadMetdata)

		if uploadResponse.Status != http.StatusOK {
			t.Errorf("Expected status to be 200, got %d", uploadResponse.Status)
		}

		if uploadResponse.Message != "request successful" {
			t.Errorf("Expected message to be request successful, got %s", uploadResponse.Message)
		}

		if uploadResponse.Item.UploadID != uploadMetdata.UploadID {
			t.Errorf("Expected UploadID to be %s, got %s", uploadMetdata.UploadID, uploadResponse.Item.UploadID)
		}
		if uploadResponse.Item.UploadedFileName != uploadMetdata.FileName {
			t.Errorf("Expected UploadedFileName to be %s, got %s", uploadMetdata.FileName, uploadResponse.Item.UploadedFileName)
		}
		if uploadResponse.Item.ProcessingStatus != "New Upload" {
			t.Errorf("Expected ProcessingStatus to be New Upload, got %s", uploadResponse.Item.ProcessingStatus)
		}

	},
	)
}

func TestUploadDetailSetStatusText(t *testing.T) {
	t.Run("Test setStatusText() method", func(t *testing.T) {
		uploadDetail := &uploadDetailDTO{}

		// UploadStatusUploaded
		uploadMetdata := &domain.UploadMetadata{
			ProcessedStatus: domain.UploadStatusUploaded,
		}
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "New Upload" {
			t.Errorf("Expected ProcessingStatus to be New Upload, got %s", uploadDetail.ProcessingStatus)
		}

		// UploadStatusProcessing
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessing
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processing" {
			t.Errorf("Expected ProcessingStatus to be Processing, got %s", uploadDetail.ProcessingStatus)
		}

		// UploadStatusProcessingAborted
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessingAborted
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processing Aborted" {
			t.Errorf("Expected ProcessingStatus to be Processing Aborted, got %s", uploadDetail.ProcessingStatus)
		}

		// UploadStatusProcessed
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessed
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processed" {
			t.Errorf("Expected ProcessingStatus to be Processed, got %s", uploadDetail.ProcessingStatus)
		}

		// Unknown status
		uploadMetdata.ProcessedStatus = 77
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Unknown Status" {
			t.Errorf("Expected ProcessingStatus to be Unknown Status, got %s", uploadDetail.ProcessingStatus)
		}
	},
	)

}
