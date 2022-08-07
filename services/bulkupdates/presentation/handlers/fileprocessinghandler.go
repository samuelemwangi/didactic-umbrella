package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/upload"
)

type FileProcessingHandler struct {
	CountryService country.CountryService
	ProductService product.ProductService
	StockService   stock.StockService
	UploadService  upload.UploadService
}

func NewFileProcessingHandler(services *application.Services) *FileProcessingHandler {
	return &FileProcessingHandler{
		CountryService: services.CountryService,
		ProductService: services.ProductService,
		StockService:   services.StockService,
		UploadService:  services.UploadService,
	}
}

func (handler *FileProcessingHandler) ProcessFile(c *gin.Context) {
	fileId := c.Param("fileid")

	filepath := ensureUploadDirectoryExists() + "/" + fileId + ".csv"

	if err := handler.UploadService.ProcessUpload(filepath, fileId); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusAccepted, "Hello Tester")

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
