package errorhelper

type ErrorDetail struct {
	ErrorMessage string `json:"errorMessage"`
}

type ErrorResponseDTO struct {
	Status  int          `json:"responseStatus"`
	Message string       `json:"responseMessage"`
	Error   *ErrorDetail `json:"errorDetails"`
}
