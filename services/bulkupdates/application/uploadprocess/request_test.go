package uploadprocess

import "testing"

func TestUploadProcessRequestDTOToEntity(t *testing.T) {
	t.Run("Test toEntity() method", func(t *testing.T) {
		uploadProcessRequest := UploadProcessRequestDTO{
			UploadID: "12356",
		}

		uploadMetada := uploadProcessRequest.toEntity()

		if uploadMetada.UploadID != uploadProcessRequest.UploadID {
			t.Errorf("Expected UploadID to be %s, got %s", uploadProcessRequest.UploadID, uploadMetada.UploadID)
		}
	},
	)
}

func TestUploadProcessRequestDTOValidateRequest(t *testing.T) {
	// Invalid request
	t.Run("Test ValidateRequest() method - Invalid Request", func(t *testing.T) {
		uploadProcessRequest := UploadProcessRequestDTO{}

		errors := uploadProcessRequest.validateRequest()

		if len(errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(errors))
		}

		if errors["UploadID"] != "required" {
			t.Errorf("Expected error message to be 'required', got %s", errors["uploadID"])
		}
	})

	// Valid request
	t.Run("Test ValidateRequest() method - Valid Request", func(t *testing.T) {
		uploadProcessRequest := UploadProcessRequestDTO{
			UploadID: "12356",
		}

		errors := uploadProcessRequest.validateRequest()

		if len(errors) != 0 {
			t.Errorf("Expected 0 error, got %d", len(errors))
		}
	})
}
