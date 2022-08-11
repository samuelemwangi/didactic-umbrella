package uploadmetadata

import (
	"testing"

	"github.com/google/uuid"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

func TestUploadMetadataDTOtoEntity(t *testing.T) {
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
	})
}
