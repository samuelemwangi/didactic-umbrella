package upload

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type uploadDetailDTO struct {
	UploadId string `json:"uploadId"`
	Status   string `json:"status"`
}

type UploadResponseDTO struct {
	Status  int              `json:"responseStatus"`
	Message string           `json:"responseMessage"`
	Item    *uploadDetailDTO `json:"itemDetails"`
}

func (response *UploadResponseDTO) toResponseDTO(uploadMetadata *domain.FileUploadMetadata) {
	uploadDetail := &uploadDetailDTO{
		UploadId: uploadMetadata.UploadId,
	}

	uploadDetail.setStatusText(uploadMetadata)

	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = uploadDetail
}

func (upload *uploadDetailDTO) setStatusText(uploadMetadata *domain.FileUploadMetadata) {

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
	default:
		upload.Status = "Unknown Status"
	}
}
