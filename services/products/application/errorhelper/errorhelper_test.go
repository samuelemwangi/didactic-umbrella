package errorhelper

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

// ============================= Test service.go

func TestErrorServiceGetValidationError(t *testing.T) {
	t.Run("Test GetValidationError() method", func(t *testing.T) {
		errorService := NewErrorService()

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
	},
	)
}

func TestErrorServiceGetGeneralError(t *testing.T) {
	t.Run("Test GetGeneralError() method", func(t *testing.T) {
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
	},
	)
}

func TestErrorServiceGetGeneralErrorWithNilError(t *testing.T) {
	t.Run("Test GetGeneralError() method with nil error", func(t *testing.T) {
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
	},
	)
}
