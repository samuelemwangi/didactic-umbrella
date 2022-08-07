package uploadprocess

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type uploadProcessDetailDTO struct {
	UploadID         string `json:"uploadId"`
	UploadedFileName string `json:"uploadedFileName"`
	ProcessingStatus string `json:"processingStatus"`
	TotalItems       uint   `json:"totalItems"`
	CreatedAt        string `json:"uploadedAt"`
}

type UploadProcessResponseDTO struct {
	Status  int                     `json:"responseStatus"`
	Message string                  `json:"responseMessage"`
	Item    *uploadProcessDetailDTO `json:"itemDetails"`
}

func (response *UploadProcessResponseDTO) toResponseDTO(entity *domain.UploadMetadata) {

	uploadProcessDetails := &uploadProcessDetailDTO{
		UploadID:         entity.UploadID,
		UploadedFileName: entity.FileName,
		TotalItems:       entity.TotalItems,
		CreatedAt:        entity.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	uploadProcessDetails.setStatusText(entity)
	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = uploadProcessDetails
}

func (upload *uploadProcessDetailDTO) setStatusText(uploadMetadata *domain.UploadMetadata) {

	switch uploadMetadata.ProcessedStatus {
	case domain.UploadStatusUploaded:
		upload.ProcessingStatus = "New"
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
