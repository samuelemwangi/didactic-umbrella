package upload

import (
	"errors"
	"strconv"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/infrastructure/fileutils"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
)

type UploadService interface {
	ProcessUpload(string, string) error
	ProcessFileData([][]string) error
}

type uploadService struct {
	uploadRepo     repositories.UploadRepository
	csvReader      fileutils.CSVReader
	countryService country.CountryService
	productService product.ProductService
	stockService   stock.StockService
	uploadId       string
}

func NewUploadService(repos *persistence.Repositories) *uploadService {
	return &uploadService{
		uploadRepo:     repos.UploadRepo,
		csvReader:      fileutils.NewCSVReader(),
		countryService: country.NewCountryService(repos),
		productService: product.NewProductService(repos),
		stockService:   stock.NewStockService(repos),
	}
}

func (service *uploadService) ProcessUpload(filePath, uploadId string) error {
	service.uploadId = uploadId

	data, err := service.csvReader.ReadFile(filePath)
	service.ManageUpdateUploadStatus(domain.UploadStatusProcessing, uint(len(data)), 0)
	if err != nil {
		return err
	}

	return service.ProcessFileData(data)
}

func (service *uploadService) ProcessFileData(data [][]string) error {

	processingErrors := make(map[string]error)
	countRecords := 0

	for i, line := range data {
		// do not process the first line
		if i > 0 {
			// save country data
			countryData, countryError := service.countryService.SaveCountry(line[0])
			if countryError != nil {
				processingErrors["country"] = countryError
				break
			}

			// save product data
			productRequest := &product.ProductRequestDTO{
				Name: line[2],
				SKU:  line[1],
			}
			productData, productError := service.productService.SaveProduct(productRequest)
			if productError != nil {
				processingErrors["country"] = productError
				break
			}

			// save stock data
			countStock, err := strconv.Atoi(line[3])
			if err != nil {
				processingErrors["stock"] = err
				break
			}
			stockRequest := &stock.StockRequestDTO{
				CountryId: countryData.CountryID,
				ProductId: productData.ProductId,
				Quantity:  countStock,
			}

			stockError := service.stockService.SaveStock(stockRequest)

			if stockError != nil {
				processingErrors["stock"] = stockError
				break
			}

			countRecords++

			break
		}
	}

	if len(processingErrors) > 0 {
		service.ManageUpdateUploadStatus(domain.UploadStatusProcessingAborted, uint(len(data)), uint(countRecords))
		return errors.New(" data processing failed")
	} else {
		service.ManageUpdateUploadStatus(domain.UploadStatusProcessingAborted, uint(len(data)), uint(countRecords))
		return nil
	}

}

func (service *uploadService) ManageUpdateUploadStatus(status uint, total uint, processed uint) error {
	uploadMetadata := &domain.FileUploadMetadata{
		UploadId:           service.uploadId,
		ProcessedStatus:    status,
		TotalItems:         total,
		TotalItemsProcesed: processed,
	}

	err := service.uploadRepo.UpdateUploadStatus(uploadMetadata)

	if err != nil {
		return err
	}

	return nil
}
