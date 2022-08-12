package uploadprocess

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/country_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/product_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/application_mock/stock_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/infrastructure_mock/fileutils_mock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock"
)

func TestGetProcessingStatus(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)
	mockCSVReader := fileutils_mock.NewMockCSVReader(mockCtrl)
	mockCountryService := country_mock.NewMockCountryService(mockCtrl)
	mockProductService := product_mock.NewMockProductService(mockCtrl)
	mockStockService := stock_mock.NewMockStockService(mockCtrl)

	uploadProcessorService := &uploadProcessorService{
		uploadMetdataRepo: mockRepos.UploadMetadataRepo,
		csvReader:         mockCSVReader,
		countryService:    mockCountryService,
		productService:    mockProductService,
		stockService:      mockStockService,
		errorService:      errorhelper.NewErrorService(),
		uploadMetadata:    &domain.UploadMetadata{},
	}

	uploadId := "sample1234"
	fileName := "sample.csv"
	totalItems := uint(10)
	totalItemsProcessed := uint(5)

	t.Run("Test GetProcessingStatus() method - valid request returns a valid response", func(t *testing.T) {

		uploadMetadata := &domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         totalItems,
			TotalItemsProcesed: totalItemsProcessed,
		}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, nil)

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if errResponse != nil {
			t.Errorf("Expected nil error, but got %v", errResponse)
		}

		if uploadResponse != nil && uploadResponse.Status != http.StatusOK {
			t.Errorf("Expected status code %v, but got %v", http.StatusOK, uploadResponse.Status)
		}

		if uploadResponse != nil && uploadResponse.Item.UploadID != uploadId {
			t.Errorf("Expected upload id %v, but got %v", uploadId, uploadResponse.Item.UploadID)
		}
	})

	t.Run("Test GetProcessingStatus() method - invalid request returns an error response", func(t *testing.T) {

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: "",
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if uploadResponse != nil && uploadResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected status code %v, but got %v", http.StatusBadRequest, uploadResponse.Status)
		}
	})

	t.Run("Test GetProcessingStatus() method - record not found error returns an error response", func(t *testing.T) {
		uploadMetadata := &domain.UploadMetadata{}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, gorm.ErrRecordNotFound)

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if uploadResponse != nil {
			t.Errorf("Expected nil response, but got %v", uploadResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusNotFound {
			t.Errorf("Expected status code %v, but got %v", http.StatusNotFound, errResponse.Status)
		}
	})

	t.Run("Test GetProcessingStatus() method - db error returns an error response", func(t *testing.T) {
		uploadMetadata := &domain.UploadMetadata{}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(uploadMetadata, errors.New("db error"))

		uploadRequestDTO := &UploadProcessRequestDTO{
			UploadID: uploadId,
		}

		uploadResponse, errResponse := uploadProcessorService.GetProcessingStatus(uploadRequestDTO)

		if uploadResponse != nil {
			t.Errorf("Expected nil response, but got %v", uploadResponse)
		}

		if errResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

		if errResponse != nil && errResponse.Status != http.StatusInternalServerError {
			t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, errResponse.Status)
		}
	})

}

func TestProcessUpload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepos := persistence_mock.NewMockRepositories(mockCtrl)
	mockCSVReader := fileutils_mock.NewMockCSVReader(mockCtrl)
	mockCountryService := country_mock.NewMockCountryService(mockCtrl)
	mockProductService := product_mock.NewMockProductService(mockCtrl)
	mockStockService := stock_mock.NewMockStockService(mockCtrl)

	uploadProcessorService := &uploadProcessorService{
		uploadMetdataRepo: mockRepos.UploadMetadataRepo,
		csvReader:         mockCSVReader,
		countryService:    mockCountryService,
		productService:    mockProductService,
		stockService:      mockStockService,
		errorService:      errorhelper.NewErrorService(),
		uploadMetadata:    &domain.UploadMetadata{},
	}

	uploadId := "sample1234"
	filepath := "/filePath/sample.csv"
	fileName := "sample.csv"

	t.Run("Test ProcessUpload() method - valid request returns a valid response", func(t *testing.T) {

		// mock get upload by upload id
		uploadMetadata := domain.UploadMetadata{
			UploadID:        uploadId,
			ProcessedStatus: domain.UploadStatusUploaded,
			FileName:        fileName,
		}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(&uploadMetadata, nil)

		// assign values

		// mock read csv data
		csvData := [][]string{
			{"country", "sku", "name", "stock_change"},
			{"eg", "9befa247cd11", "Chung PLC Table", "22"},
		}

		mockCSVReader.EXPECT().ReadFile(filepath).Return(csvData, nil)

		// mock update processing
		uploadMetadataProcessing := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}

		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessing).Return(nil)

		// mock save Country data
		country := &country.CountryItemDTO{
			CountryID:   1,
			CountryCode: csvData[1][0],
		}

		mockCountryService.EXPECT().SaveCountry(csvData[1][0]).Return(country, nil)

		// mock save product Data
		product := &product.ProductItemDTO{
			ProductID:   1,
			SKU:         csvData[1][1],
			ProductName: csvData[1][2],
		}
		mockProductService.EXPECT().SaveProduct(csvData[1][1], csvData[1][2]).Return(product, nil)

		// mock save stock change data
		countStock, _ := strconv.Atoi(csvData[1][3])
		mockStockService.EXPECT().SaveStock(country.CountryID, product.ProductID, countStock).Return(nil)

		// update status to processed
		uploadMetadataProcessed := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessed,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: uint(len(csvData) - 1),
		}

		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessed).Return(nil)

		// invoke process upload
		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse != nil {
			t.Errorf("Expected nil error, but got %v", errorResponse)
		}

	})

	t.Run("Test ProcessUpload() method - db error when getting an upload returns an error response", func(t *testing.T) {
		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(nil, errors.New("db error"))
		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test ProcessUpload() method - processed upload returns an error response", func(t *testing.T) {
		uploadMetadata := domain.UploadMetadata{
			UploadID:        uploadId,
			ProcessedStatus: domain.UploadStatusProcessed,
			FileName:        fileName,
		}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(&uploadMetadata, nil)

		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test ProcessUpload() method - service error when saving country returns an error response", func(t *testing.T) {
		// mock get upload by upload id
		uploadMetadata := domain.UploadMetadata{
			UploadID:        uploadId,
			ProcessedStatus: domain.UploadStatusUploaded,
			FileName:        fileName,
		}
		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(&uploadMetadata, nil)

		// mock read csv data
		csvData := [][]string{
			{"country", "sku", "name", "stock_change"},
			{"eg", "9befa247cd11", "Chung PLC Table", "22"},
		}

		mockCSVReader.EXPECT().ReadFile(filepath).Return(csvData, nil)

		// mock update processing
		uploadMetadataProcessing := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}
		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessing).Return(nil)

		// mock save Country data
		mockCountryService.EXPECT().SaveCountry(csvData[1][0]).Return(nil, errors.New("service error"))

		// mock update processing for aborted update
		uploadMetadataProcessingAborted := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessingAborted,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}
		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessingAborted).Return(nil)

		// Send request
		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test ProcessUpload() method - service error when saving product returns an error response", func(t *testing.T) {
		// mock get upload by upload id
		uploadMetadata := domain.UploadMetadata{
			UploadID:        uploadId,
			ProcessedStatus: domain.UploadStatusUploaded,
			FileName:        fileName,
		}
		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(&uploadMetadata, nil)

		// mock read csv data
		csvData := [][]string{
			{"country", "sku", "name", "stock_change"},
			{"eg", "9befa247cd11", "Chung PLC Table", "22"},
		}

		mockCSVReader.EXPECT().ReadFile(filepath).Return(csvData, nil)

		// mock update processing
		uploadMetadataProcessing := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}
		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessing).Return(nil)

		// mock save Country data
		// mock save Country data
		country := &country.CountryItemDTO{
			CountryID:   1,
			CountryCode: csvData[1][0],
		}

		mockCountryService.EXPECT().SaveCountry(csvData[1][0]).Return(country, nil)

		// mock save Product data
		mockProductService.EXPECT().SaveProduct(csvData[1][1], csvData[1][2]).Return(nil, errors.New("service error"))

		// mock update processing for aborted update
		uploadMetadataProcessingAborted := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessingAborted,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}
		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessingAborted).Return(nil)

		// Send request
		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("Test ProcessUpload() method - service error when saving stock returns an error response", func(t *testing.T) {

		// mock get upload by upload id
		uploadMetadata := domain.UploadMetadata{
			UploadID:        uploadId,
			ProcessedStatus: domain.UploadStatusUploaded,
			FileName:        fileName,
		}

		mockRepos.UploadMetadataRepo.EXPECT().GetUploadByUploadId(uploadId).Return(&uploadMetadata, nil)

		// assign values

		// mock read csv data
		csvData := [][]string{
			{"country", "sku", "name", "stock_change"},
			{"eg", "9befa247cd11", "Chung PLC Table", "22"},
		}

		mockCSVReader.EXPECT().ReadFile(filepath).Return(csvData, nil)

		// mock update processing
		uploadMetadataProcessing := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessing,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}

		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessing).Return(nil)

		// mock save Country data
		country := &country.CountryItemDTO{
			CountryID:   1,
			CountryCode: csvData[1][0],
		}

		mockCountryService.EXPECT().SaveCountry(csvData[1][0]).Return(country, nil)

		// mock save product Data
		product := &product.ProductItemDTO{
			ProductID:   1,
			SKU:         csvData[1][1],
			ProductName: csvData[1][2],
		}
		mockProductService.EXPECT().SaveProduct(csvData[1][1], csvData[1][2]).Return(product, nil)

		// mock save stock change data
		countStock, _ := strconv.Atoi(csvData[1][3])
		mockStockService.EXPECT().SaveStock(country.CountryID, product.ProductID, countStock).Return(errors.New("service error"))

		// update status to processed
		uploadMetadataProcessingAborted := domain.UploadMetadata{
			UploadID:           uploadId,
			ProcessedStatus:    domain.UploadStatusProcessingAborted,
			FileName:           fileName,
			TotalItems:         uint(len(csvData) - 1),
			TotalItemsProcesed: 0,
		}

		mockRepos.UploadMetadataRepo.EXPECT().UpdateUploadStatus(&uploadMetadataProcessingAborted).Return(nil)

		// invoke process upload
		errorResponse := uploadProcessorService.ProcessUpload(filepath, uploadId)

		if errorResponse == nil {
			t.Errorf("Expected error, but got nil")
		}

	})

}
