package upload

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type UploadService interface {
	SaveUploadMetadaData(*UploadMetadataDTO) (*UploadResponseDTO, *error.ErrorResponseDTO)
}

type uploadService struct {
	uploadRepo   repositories.UploadRepository
	errorService error.ErrorService
}

func NewUploadService(repos *persistence.Repositories) *uploadService {
	return &uploadService{
		uploadRepo:   repos.UploadRepo,
		errorService: error.NewErrorService(),
	}
}

func (service *uploadService) SaveUploadMetadaData(request *UploadMetadataDTO) (*UploadResponseDTO, *error.ErrorResponseDTO) {

	data, dbError := service.uploadRepo.SaveUploadMetaData(request.toEntity())

	if dbError != nil {
		return nil, service.errorService.GetGeneralError(http.StatusInternalServerError, *dbError)
	}

	response := UploadResponseDTO{}
	response.toResponseDTO(data)

	return &response, nil
}
