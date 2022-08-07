package errorhelper

import "strings"

type ErrorService interface {
	GetValidationError(int, map[string]string) *ErrorResponseDTO
	GetGeneralError(int, error) *ErrorResponseDTO
}

type errorService struct {
}

func NewErrorService() *errorService {
	return &errorService{}
}

func (es *errorService) GetValidationError(status int, validationErrors map[string]string) *ErrorResponseDTO {

	errorMessage := ""

	for key, value := range validationErrors {
		errorMessage += key + ": " + value + ","
	}

	errorDetails := &ErrorDetail{
		ErrorMessage: strings.TrimRight(errorMessage, ","),
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: "request validation failed",
		Error:   errorDetails,
	}
}

func (es *errorService) GetGeneralError(status int, err error) *ErrorResponseDTO {
	errorDetails := &ErrorDetail{
		ErrorMessage: err.Error(),
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: "request failed",
		Error:   errorDetails,
	}
}
