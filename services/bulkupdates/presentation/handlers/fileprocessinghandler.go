package handlers

import (
	"log"

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
	fileId := c.Param("fileId")

	uploadRequest := &upload.UploadMetadataRequestDTO{
		UploadId: fileId,
	}

	if err := handler.UploadService.SaveUpload(uploadRequest); err != nil {
		log.Println(err)
	}

}
