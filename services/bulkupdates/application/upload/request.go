package upload

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type UploadMetadataRequestDTO struct {
	UploadId string
	Status   uint
}

func (request *UploadMetadataRequestDTO) toEntity() *domain.FileUploadMetadata {
	return &domain.FileUploadMetadata{
		UploadId:        request.UploadId,
		ProcessedStatus: request.Status,
	}
}
