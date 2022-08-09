package errorhelper

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

// ============================= Test service.go

func TestErrorServiceGetValidationError(t *testing.T) {
	errorService := NewErrorService()

	// Single error
	t.Run("Test GetValidationError() method - Single Validation Error", func(t *testing.T) {

		validationErrors := map[string]string{
			"name": "name is required",
		}
		errorResponse := errorService.GetValidationError(http.StatusBadRequest, validationErrors)

		if errorResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status to be 400, got %d", errorResponse.Status)
		}

		if errorResponse.Message != "request validation failed" {
			t.Errorf("Expected message to be request validation failed, got %s", errorResponse.Message)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "name is required") {
			t.Errorf("Expected error message to contain 'name is required', got %s", errorResponse.Error.ErrorMessage)
		}
	})

	// Multiple errors
	t.Run("Test GetValidationError() method - Multiple Validation Errors", func(t *testing.T) {

		validationErrors := map[string]string{
			"name": "name is required",
			"age":  "age is required",
		}
		errorResponse := errorService.GetValidationError(http.StatusBadRequest, validationErrors)

		if errorResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status to be 400, got %d", errorResponse.Status)
		}

		if errorResponse.Message != "request validation failed" {
			t.Errorf("Expected message to be request validation failed, got %s", errorResponse.Message)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "name is required") {
			t.Errorf("Expected error message to contain 'name is required', got %s", errorResponse.Error.ErrorMessage)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "age is required") {
			t.Errorf("Expected error message to contain 'age is required', got %s", errorResponse.Error.ErrorMessage)
		}
	})

	// No errors
	t.Run("Test GetValidationError() method - No Validation Errors", func(t *testing.T) {

		errorResponse := errorService.GetValidationError(http.StatusBadRequest, nil)

		if errorResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status to be 400, got %d", errorResponse.Status)
		}

		if errorResponse.Message != "request validation failed" {
			t.Errorf("Expected message to be request validation failed, got %s", errorResponse.Message)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "kindly check your request and try again") {
			t.Errorf("Expected error message to contain 'kindly check your request and try again', got %s", errorResponse.Error.ErrorMessage)
		}
	})
}

func TestErrorServiceGetGeneralError(t *testing.T) {
	// Non-nil error
	t.Run("Test GetGeneralError() method - Non-nil error", func(t *testing.T) {
		errorService := NewErrorService()

		err := errors.New("error message")

		errorResponse := errorService.GetGeneralError(http.StatusInternalServerError, err)

		if errorResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status to be 500, got %d", errorResponse.Status)
		}

		if errorResponse.Message != "request failed" {
			t.Errorf("Expected message to be request failed, got %s", errorResponse.Message)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "error message") {
			t.Errorf("Expected error message to contain 'error message', got %s", errorResponse.Error.ErrorMessage)
		}
	})

	// nil error
	t.Run("Test GetGeneralError() method - nil error", func(t *testing.T) {
		errorService := NewErrorService()

		errorResponse := errorService.GetGeneralError(http.StatusInternalServerError, nil)

		if errorResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status to be 500, got %d", errorResponse.Status)
		}

		if errorResponse.Message != "request failed" {
			t.Errorf("Expected message to be request failed, got %s", errorResponse.Message)
		}

		if !strings.Contains(errorResponse.Error.ErrorMessage, "a system error has occurred") {
			t.Errorf("Expected error message to contain 'a system error has occurred', got %s", errorResponse.Error.ErrorMessage)
		}
	})
}
