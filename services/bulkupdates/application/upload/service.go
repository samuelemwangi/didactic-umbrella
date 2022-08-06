package upload

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
)

type UploadService interface {
	ProcessUpload(string) error
}

type uploadService struct {
	uploadRepo repositories.UploadRepository
}

func NewUploadService(repos *persistence.Repositories) *uploadService {
	return &uploadService{
		uploadRepo: repos.UploadRepo,
	}
}

func (service *uploadService) ProcessUpload(uploadId string) error {
	uploadMetadata := &domain.FileUploadMetadata{
		UploadId:        uploadId,
		ProcessedStatus: domain.UploadStatusProcessing,
	}

	err := service.uploadRepo.UpdateUploadStatus(uploadMetadata)

	if err != nil {
		return err
	}

	return nil
}
