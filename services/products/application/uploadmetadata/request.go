package uploadmetadata

import "github.com/samuelemwangi/jumia-mds-test/services/products/domain"

type UploadMetadataDTO struct {
	FileName string
	UploadID string
}

func (request *UploadMetadataDTO) toEntity() *domain.UploadMetadata {

	return &domain.UploadMetadata{
		FileName:        request.FileName,
		UploadID:        request.UploadID,
		ProcessedStatus: domain.UploadStatusUploaded,
	}
}
