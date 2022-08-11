package uploadprocess

import (
	"testing"
	"time"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

func TestUploadProcesResponseToDTO(t *testing.T) {
	t.Run("Test toResponseDTO() method", func(t *testing.T) {

		uploadMetadata := &domain.UploadMetadata{
			UploadID:        "12356",
			FileName:        "test.csv",
			TotalItems:      10,
			ProcessedStatus: domain.UploadStatusProcessed,
		}
		uploadMetadata.CreatedAt = time.Now()

		uploadProcessResponse := UploadProcessResponseDTO{}
		uploadProcessResponse.toResponseDTO(uploadMetadata)

		if uploadProcessResponse.Item.UploadID != uploadMetadata.UploadID {
			t.Errorf("Expected UploadID to be %s, got %s", uploadMetadata.UploadID, uploadProcessResponse.Item.UploadID)
		}

		if uploadProcessResponse.Item.UploadedFileName != uploadMetadata.FileName {
			t.Errorf("Expected UploadedFileName to be %s, got %s", uploadMetadata.FileName, uploadProcessResponse.Item.UploadedFileName)
		}

		if uploadProcessResponse.Item.TotalItems != uploadMetadata.TotalItems {
			t.Errorf("Expected TotalItems to be %d, got %d", uploadMetadata.TotalItems, uploadProcessResponse.Item.TotalItems)
		}

		if uploadProcessResponse.Item.CreatedAt != uploadMetadata.CreatedAt.Format("2006-01-02 15:04:05") {
			t.Errorf("Expected CreatedAt to be %s, got %s", uploadMetadata.CreatedAt.Format("2006-01-02 15:04:05"), uploadProcessResponse.Item.CreatedAt)
		}

		if uploadProcessResponse.Item.ProcessingStatus != "Processed" {
			t.Errorf("Expected ProcessedStatus to be 'Processed', got %s", uploadProcessResponse.Item.ProcessingStatus)
		}

	})
}

func TestUploadDetailSetStatusText(t *testing.T) {
	uploadDetail := &uploadProcessDetailDTO{}

	// UploadStatusUploaded
	uploadMetdata := &domain.UploadMetadata{
		ProcessedStatus: domain.UploadStatusUploaded,
	}
	uploadDetail.setStatusText(uploadMetdata)

	t.Run("Test setStatusText() method - UploadStatusUploaded", func(t *testing.T) {

		if uploadDetail.ProcessingStatus != "New Upload" {
			t.Errorf("Expected ProcessingStatus to be New Upload, got %s", uploadDetail.ProcessingStatus)
		}

	})

	t.Run("Test setStatusText() method - UploadStatusProcessing", func(t *testing.T) {
		// UploadStatusProcessing
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessing
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processing" {
			t.Errorf("Expected ProcessingStatus to be Processing, got %s", uploadDetail.ProcessingStatus)
		}
	})

	t.Run("Test setStatusText() method - UploadStatusProcessingAborted", func(t *testing.T) {
		// UploadStatusProcessingAborted
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessingAborted
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processing Aborted" {
			t.Errorf("Expected ProcessingStatus to be Processing Aborted, got %s", uploadDetail.ProcessingStatus)
		}
	})

	t.Run("Test setStatusText() method - UploadStatusProcessed", func(t *testing.T) {
		// UploadStatusProcessed
		uploadMetdata.ProcessedStatus = domain.UploadStatusProcessed
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Processed" {
			t.Errorf("Expected ProcessingStatus to be Processed, got %s", uploadDetail.ProcessingStatus)
		}
	})

	t.Run("Test setStatusText() method - Unknown Status", func(t *testing.T) {
		// Unknown status
		uploadMetdata.ProcessedStatus = 77
		uploadDetail.setStatusText(uploadMetdata)

		if uploadDetail.ProcessingStatus != "Unknown Status" {
			t.Errorf("Expected ProcessingStatus to be Unknown Status, got %s", uploadDetail.ProcessingStatus)
		}
	})

}
