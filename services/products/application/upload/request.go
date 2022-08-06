package upload

import "github.com/samuelemwangi/jumia-mds-test/services/products/domain"

type UploadMetadataDTO struct {
	FileName string
	UploadId string
}

func (request *UploadMetadataDTO) toEntity() *domain.FileUploadMetadata {

	return &domain.FileUploadMetadata{
		FileName:        request.FileName,
		UploadId:        request.UploadId,
		ProcessedStatus: domain.UploadStatusUploaded,
	}
}
