package uploadmetadata

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/infrastructure/queueing"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type UploadMetadataService interface {
	SaveUploadMetadaData(*UploadMetadataDTO) (*UploadResponseDTO, *errorhelper.ErrorResponseDTO)
}

type uploadMetadataService struct {
	uploadMetadataRepo repositories.UploadMetadataRepository
	errorService       errorhelper.ErrorService
	kafkaProducer      queueing.KafkaProducer
	fileIdTopic        string
}

func NewUploadMetadataService(repos *persistence.Repositories) *uploadMetadataService {
	return &uploadMetadataService{
		uploadMetadataRepo: repos.UploadMetadataRepo,
		errorService:       errorhelper.NewErrorService(),
		kafkaProducer:      queueing.NewKafkaProducer(),
		fileIdTopic:        "file-process-topic",
	}
}

func (service *uploadMetadataService) SaveUploadMetadaData(request *UploadMetadataDTO) (*UploadResponseDTO, *errorhelper.ErrorResponseDTO) {

	uploadMetadata := request.toEntity()

	dbError := service.uploadMetadataRepo.SaveUploadMetaData(uploadMetadata)

	if dbError != nil {
		return nil, service.errorService.GetGeneralError(http.StatusInternalServerError, dbError)
	}

	response := UploadResponseDTO{}
	response.toUploadResponseDTO(uploadMetadata)

	service.kafkaProducer.ProduceMessage(service.fileIdTopic, request.UploadID)

	return &response, nil
}
