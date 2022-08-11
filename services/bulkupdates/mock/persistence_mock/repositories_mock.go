package persistence_mock

import (
	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/mock/persistence_mock/repositories_mock"
)

type MockRepositories struct {
	CountryRepo        *repositories_mock.MockCountryRepository
	ProductRepo        *repositories_mock.MockProductRepository
	StockRepo          *repositories_mock.MockStockRepository
	UploadMetadataRepo *repositories_mock.MockUploadMetadataRepository
}

func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	return &MockRepositories{
		CountryRepo:        repositories_mock.NewMockCountryRepository(ctrl),
		ProductRepo:        repositories_mock.NewMockProductRepository(ctrl),
		StockRepo:          repositories_mock.NewMockStockRepository(ctrl),
		UploadMetadataRepo: repositories_mock.NewMockUploadMetadataRepository(ctrl),
	}

}
