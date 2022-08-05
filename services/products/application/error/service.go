package error

import "strings"

type ErrorService interface {
	GetValidationError(int, string, map[string]string) *ErrorResponseDTO
	GetGeneralError(int, string) *ErrorResponseDTO
}

type errorService struct {
}

func NewErrorService() *errorService {
	return &errorService{}
}

func (es *errorService) GetValidationError(status int, message string, validationErrors map[string]string) *ErrorResponseDTO {

	errorMessage := ""

	for key, value := range validationErrors {
		errorMessage += key + ": " + value + ","
	}

	errorDetails := &ErrorDetail{
		ErrorMessage: strings.TrimRight(errorMessage, ","),
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: message,
		Error:   errorDetails,
	}
}

func (es *errorService) GetGeneralError(status int, message string) *ErrorResponseDTO {
	errorDetails := &ErrorDetail{
		ErrorMessage: message,
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: "request failed",
		Error:   errorDetails,
	}
}
