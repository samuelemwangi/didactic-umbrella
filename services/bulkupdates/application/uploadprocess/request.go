package uploadprocess

import (
	"github.com/go-playground/validator/v10"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadProcessRequestDTO struct {
	UploadID string
}

func (request *UploadProcessRequestDTO) toEntity() *domain.UploadMetadata {
	return &domain.UploadMetadata{
		UploadID: request.UploadID,
	}
}

func (request *UploadProcessRequestDTO) validateRequest() map[string]string {
	errors := make(map[string]string)

	err := validator.New().Struct(request)
	if err == nil {
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.ActualTag()
	}

	return errors
}
