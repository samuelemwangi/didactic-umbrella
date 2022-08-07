package handlers

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/uploadmetadata"
)

type UploadHandler struct {
	uploadMetadataService uploadmetadata.UploadMetadataService
	errorService          errorhelper.ErrorService
}

func NewUploadHandler(services *application.Services) *UploadHandler {
	return &UploadHandler{
		uploadMetadataService: services.UploadMetadataService,
		errorService:          services.ErrorService,
	}
}

func (handler *UploadHandler) UploadCSVFile(c *gin.Context) {
	file, err := c.FormFile("file")

	// validate inputs
	if err != nil {
		errorResponse := handler.errorService.GetGeneralError(http.StatusBadRequest, err)
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	extension := filepath.Ext(file.Filename)

	if strings.ToLower(extension) != ".csv" {
		err = errors.New("file not supported. kindly upload a csv file")
		errorResponse := handler.errorService.GetGeneralError(http.StatusBadRequest, err)
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	// upload file
	uploadID := uuid.New().String()

	uploadPath := ensureUploadDirectoryExists()

	if err := c.SaveUploadedFile(file, uploadPath+"/"+uploadID+extension); err != nil {
		err = errors.New("file upload failed. kindly retry")
		errorResponse := handler.errorService.GetGeneralError(http.StatusInternalServerError, err)
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	uploadrequest := uploadmetadata.UploadMetadataDTO{
		FileName: file.Filename,
		UploadID: uploadID,
	}

	updateResponse, errorResponse := handler.uploadMetadataService.SaveUploadMetadaData(&uploadrequest)

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	c.JSON(updateResponse.Status, updateResponse)
}

func ensureUploadDirectoryExists() string {
	uploadPath := os.Getenv("FILE_PATH")
	if uploadPath == "" {
		uploadPath = "../../uploads/csv"
	}

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, 0755)
	}

	return uploadPath
}
