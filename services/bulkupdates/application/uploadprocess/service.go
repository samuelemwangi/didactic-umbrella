package uploadprocess

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/infrastructure/fileutils"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence/repositories"
	"gorm.io/gorm"
)

type UploadProcessorService interface {
	ProcessUpload(string, string) error
	GetProcessingStatus(*UploadProcessRequestDTO) (*UploadProcessResponseDTO, *errorhelper.ErrorResponseDTO)
}

type uploadProcessorService struct {
	uploadMetdataRepo repositories.UploadMetadataRepository
	csvReader         fileutils.CSVReader
	countryService    country.CountryService
	productService    product.ProductService
	stockService      stock.StockService
	errorService      errorhelper.ErrorService
	uploadMetadata    *domain.UploadMetadata
}

func NewUploadProcessorService(repos *persistence.Repositories) *uploadProcessorService {
	return &uploadProcessorService{
		uploadMetdataRepo: repos.UploadMetdataRepo,
		csvReader:         fileutils.NewCSVReader(),
		countryService:    country.NewCountryService(repos),
		productService:    product.NewProductService(repos),
		stockService:      stock.NewStockService(repos),
		errorService:      errorhelper.NewErrorService(),
		uploadMetadata:    &domain.UploadMetadata{},
	}
}

func (service *uploadProcessorService) GetProcessingStatus(requesst *UploadProcessRequestDTO) (*UploadProcessResponseDTO, *errorhelper.ErrorResponseDTO) {
	// validate request
	validationErrors := requesst.validateRequest()

	if len(validationErrors) > 0 {
		return nil, service.errorService.GetValidationError(http.StatusBadRequest, validationErrors)
	}

	uploadMetadata, dbError := service.uploadMetdataRepo.GetUploadByUploadId(requesst.UploadID)

	// handle errors
	if dbError != nil {
		status := http.StatusInternalServerError
		if dbError.Error() == gorm.ErrRecordNotFound.Error() {
			status = http.StatusNotFound
		}
		return nil, service.errorService.GetGeneralError(status, dbError)
	}

	// prepare response
	var responseDTO UploadProcessResponseDTO
	responseDTO.toResponseDTO(uploadMetadata)
	return &responseDTO, nil
}

func (service *uploadProcessorService) ProcessUpload(filePath, uploadId string) error {

	// check if upload has been processed
	uploadMetadata, err := service.uploadMetdataRepo.GetUploadByUploadId(uploadId)
	if err != nil {
		return err
	}

	service.uploadMetadata = uploadMetadata

	if service.uploadMetadata.ProcessedStatus == domain.UploadStatusProcessed {
		return errors.New("upload has already been processed")
	}

	// read file data
	data, err := service.csvReader.ReadFile(filePath)
	if err != nil {
		return err
	}

	// indicate we have started processing the file
	service.manageUpdateUploadStatus(domain.UploadStatusProcessing, uint(len(data)), 0)

	// process the read file data
	return service.processFileData(data)
}

func (service *uploadProcessorService) processFileData(data [][]string) error {

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
		}
	}

	if processingError.Error() != "" {
		service.manageUpdateUploadStatus(domain.UploadStatusProcessingAborted, uint(len(data)), uint(countRecords))
		return processingError
	} else {
		service.manageUpdateUploadStatus(domain.UploadStatusProcessed, uint(len(data)), uint(countRecords))
		return nil
	}

}

func (service *uploadProcessorService) manageUpdateUploadStatus(status uint, total uint, processed uint) error {

	service.uploadMetadata.ProcessedStatus = status
	//  we ignore the 1st line
	service.uploadMetadata.TotalItems = total - 1
	service.uploadMetadata.TotalItemsProcesed = processed
	err := service.uploadMetdataRepo.UpdateUploadStatus(service.uploadMetadata)

	if err != nil {
		return err
	}

	return nil
}
