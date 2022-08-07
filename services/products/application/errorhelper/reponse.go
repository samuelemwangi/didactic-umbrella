package errorhelper

type errorDetail struct {
	ErrorMessage string `json:"errorMessage"`
}

type ErrorResponseDTO struct {
	Status  int          `json:"responseStatus"`
	Message string       `json:"responseMessage"`
	Error   *errorDetail `json:"errorDetails"`
}
