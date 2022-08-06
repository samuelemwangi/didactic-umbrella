package upload

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
)

type UploadService interface {
	SaveUpload(*UploadMetadataRequestDTO) error
}

type uploadService struct {
	uploadRepo repositories.UploadRepository
}

func NewUploadService(repos *persistence.Repositories) *uploadService {
	return &uploadService{
		uploadRepo: repos.UploadRepo,
	}
}

func (service *uploadService) SaveUpload(request *UploadMetadataRequestDTO) error {
	upload := request.toEntity()
	err := service.uploadRepo.UpdateUpload(upload)

	return err
}
