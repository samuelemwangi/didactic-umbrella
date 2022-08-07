package uploadmetadata

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type uploadMetadataDetailDTO struct {
	UploadID string `json:"uploadId"`
	Status   string `json:"status"`
}

type UploadResponseDTO struct {
	Status  int                      `json:"responseStatus"`
	Message string                   `json:"responseMessage"`
	Item    *uploadMetadataDetailDTO `json:"itemDetails"`
}

func (response *UploadResponseDTO) toResponseDTO(uploadMetadata *domain.UploadMetadata) {
	uploadDetail := &uploadMetadataDetailDTO{
		UploadID: uploadMetadata.UploadID,
	}

	uploadDetail.setStatusText(uploadMetadata)

	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = uploadDetail
}

func (upload *uploadMetadataDetailDTO) setStatusText(uploadMetadata *domain.UploadMetadata) {

	switch uploadMetadata.ProcessedStatus {
	case domain.UploadStatusUploaded:
		upload.Status = "New"
		return
	case domain.UploadStatusProcessing:
		upload.Status = "Processing"
		return
	case domain.UploadStatusProcessed:
		upload.Status = "Processed"
		return
	case domain.UploadStatusProcessingAborted:
		upload.Status = "Processing Aborted"
		return
	default:
		upload.Status = "Unknown Status"
	}
}
