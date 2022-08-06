package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/upload"
)

type UploadHandler struct {
	uploadService upload.UploadService
	errorService  error.ErrorService
}

func NewUploadHandler(services *application.Services) *UploadHandler {
	return &UploadHandler{
		uploadService: services.UploadService,
		errorService:  services.ErrorService,
	}
}

func (handler *UploadHandler) UploadCSVFile(c *gin.Context) {
	file, err := c.FormFile("file")

	// validate inputs
	if err != nil {
		errorResponse := handler.errorService.GetGeneralError(http.StatusBadRequest, err.Error())
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	extension := filepath.Ext(file.Filename)

	if !strings.Contains(strings.ToLower(extension), "csv") {
		errorResponse := handler.errorService.GetGeneralError(http.StatusBadRequest, "file not supported. kindly upload a csv file")
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	// upload file
	fileId := uuid.New().String() + extension

	uploadPath := ensureUploadDirectoryExists()

	if err := c.SaveUploadedFile(file, uploadPath+fileId); err != nil {
		log.Fatalln(err.Error())
		errorResponse := handler.errorService.GetGeneralError(http.StatusInternalServerError, "file upload failed. kindly retry")
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	uploadrequest := upload.UploadMetadataDTO{
		UploadId: fileId,
	}

	updateResponse, errorResponse := handler.uploadService.SaveUploadMetadaData(&uploadrequest)

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
