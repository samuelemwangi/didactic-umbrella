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

	if validationErrors == nil {
		errorMessage = "kindly check your request and try again"
	}

	errorDetails := &errorDetail{
		ErrorMessage: strings.TrimRight(errorMessage, ","),
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: "request validation failed",
		Error:   errorDetails,
	}
}

func (es *errorService) GetGeneralError(status int, err error) *ErrorResponseDTO {
	errorMessage := "a system error has occurred"
	if err != nil {
		errorMessage = err.Error()
	}
	errorDetails := &errorDetail{
		ErrorMessage: errorMessage,
	}

	return &ErrorResponseDTO{
		Status:  status,
		Message: "request failed",
		Error:   errorDetails,
	}
}
