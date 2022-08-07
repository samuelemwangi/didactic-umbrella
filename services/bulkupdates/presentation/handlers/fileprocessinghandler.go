package handlers

import (
	"log"
	"os"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/uploadmetadata"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/infrastructure/queueing"
)

type FileProcessingHandler struct {
	CountryService        country.CountryService
	ProductService        product.ProductService
	StockService          stock.StockService
	UploadMetadataService uploadmetadata.UploadMetadataService
	kafkaConsumer         queueing.KafkaConsumer
}

func NewFileProcessingHandler(services *application.Services) *FileProcessingHandler {
	return &FileProcessingHandler{
		CountryService:        services.CountryService,
		ProductService:        services.ProductService,
		StockService:          services.StockService,
		UploadMetadataService: services.UploadMetadataService,
		kafkaConsumer:         queueing.NewKafkaConsumer(),
	}
}

func (handler *FileProcessingHandler) ProcessFile() {

	for {
		messageId, consumerError := handler.kafkaConsumer.ConsumeMessage("file-process-topic")

		if consumerError != nil {
			log.Println(consumerError)
		} else {
			log.Println()
			log.Println("Processing Data")
			log.Println("messageId: ", *messageId)
			log.Println()
			log.Println()
			filepath := ensureUploadDirectoryExists() + "/" + *messageId + ".csv"
			if err := handler.UploadMetadataService.ProcessUpload(filepath, *messageId); err != nil {
				log.Println(err)
			} else {
				log.Println()
				log.Println("Processing Completed")
				log.Println()
				log.Println()
			}
		}
	}

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
