package uploadmetadata

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

type UploadMetadataService interface {
	ProcessUpload(string, string) error
	ProcessFileData([][]string) error
}

type uploadMetadataService struct {
	uploadRepo     repositories.UploadRepository
	csvReader      fileutils.CSVReader
	countryService country.CountryService
	productService product.ProductService
	stockService   stock.StockService
	uploadMetadata *domain.UploadMetadata
}

func NewUploadMetadataService(repos *persistence.Repositories) *uploadMetadataService {
	return &uploadMetadataService{
		uploadRepo:     repos.UploadRepo,
		csvReader:      fileutils.NewCSVReader(),
		countryService: country.NewCountryService(repos),
		productService: product.NewProductService(repos),
		stockService:   stock.NewStockService(repos),
		uploadMetadata: &domain.UploadMetadata{},
	}
}

func (service *uploadMetadataService) ProcessUpload(filePath, uploadId string) error {
	// assign upload id to upload metadata
	service.uploadMetadata.UploadID = uploadId

	// check if upload has been processed
	err := service.uploadRepo.GetUploadByUploadId(service.uploadMetadata)
	if err != nil {
		return err
	}

	if service.uploadMetadata.ProcessedStatus == domain.UploadStatusProcessed {
		return errors.New("upload has already been processed")
	}

	// read file data
	data, err := service.csvReader.ReadFile(filePath)
	if err != nil {
		return err
	}

	// indicate we have started processing the file
	service.ManageUpdateUploadStatus(domain.UploadStatusProcessing, uint(len(data)), 0)

	// process the read file data
	return service.ProcessFileData(data)
}

func (service *uploadMetadataService) ProcessFileData(data [][]string) error {

	processingError := errors.New("")
	countRecords := 0

	for i, line := range data {
		// do not process the first line - assumably the header
		if i > 0 {
			// save country data
			countryData, countryError := service.countryService.SaveCountry(line[0])
			if countryError != nil {
				processingError = countryError
				break
			}

			// save product data ,pass sku and name
			productData, productError := service.productService.SaveProduct(line[1], line[2])
			if productError != nil {
				processingError = productError
				break
			}

			// save stock data
			countStock, err := strconv.Atoi(line[3])
			if err != nil {
				processingError = err
				break
			}

			stockError := service.stockService.SaveStock(countryData.CountryID, productData.ProductID, countStock)

			if stockError != nil {
				processingError = stockError
				break
			}

			countRecords++

			// comment out later to process all records
			break
		}
	}

	if processingError.Error() != "" {
		service.ManageUpdateUploadStatus(domain.UploadStatusProcessingAborted, uint(len(data)), uint(countRecords))
		return processingError
	} else {
		service.ManageUpdateUploadStatus(domain.UploadStatusProcessed, uint(len(data)), uint(countRecords))
		return nil
	}

}

func (service *uploadMetadataService) ManageUpdateUploadStatus(status uint, total uint, processed uint) error {

	service.uploadMetadata.ProcessedStatus = status
	service.uploadMetadata.TotalItems = total
	service.uploadMetadata.TotalItemsProcesed = processed
	err := service.uploadRepo.UpdateUploadStatus(service.uploadMetadata)

	if err != nil {
		return err
	}

	return nil
}