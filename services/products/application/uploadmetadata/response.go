package uploadmetadata

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type uploadDetailDTO struct {
	UploadID         string `json:"uploadId"`
	UploadedFileName string `json:"uploadedFileName"`
	ProcessingStatus string `json:"processingStatus"`
	CreatedAt        string `json:"createdAt"`
}

type UploadResponseDTO struct {
	Status  int              `json:"responseStatus"`
	Message string           `json:"responseMessage"`
	Item    *uploadDetailDTO `json:"itemDetails"`
}

func (response *UploadResponseDTO) toUploadResponseDTO(entity *domain.UploadMetadata) {

	uploadDetails := &uploadDetailDTO{
		UploadID:         entity.UploadID,
		UploadedFileName: entity.FileName,
		CreatedAt:        entity.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	uploadDetails.setStatusText(entity)
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = uploadDetails
}

func (upload *uploadDetailDTO) setStatusText(uploadMetadata *domain.UploadMetadata) {

	switch uploadMetadata.ProcessedStatus {
	case domain.UploadStatusUploaded:
		upload.ProcessingStatus = "New Upload"
		return
	case domain.UploadStatusProcessing:
		upload.ProcessingStatus = "Processing"
		return
	case domain.UploadStatusProcessed:
		upload.ProcessingStatus = "Processed"
		return
	case domain.UploadStatusProcessingAborted:
		upload.ProcessingStatus = "Processing Aborted"
		return
	default:
		upload.ProcessingStatus = "Unknown Status"
	}
}
