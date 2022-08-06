package upload

import "github.com/samuelemwangi/jumia-mds-test/services/products/domain"

type UploadMetadataDTO struct {
	UploadId string
}

func (request *UploadMetadataDTO) toEntity() *domain.FileUploadMetadata {

	return &domain.FileUploadMetadata{
		UploadId:        request.UploadId,
		ProcessedStatus: domain.UploadStatusUploaded,
	}
}
