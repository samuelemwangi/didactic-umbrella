package mock_persistence

import (
	"github.com/golang/mock/gomock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/mock/mock_persistence/mock_repositories"
)

type MockRepositories struct {
	CountryRepo        *mock_repositories.MockCountryRepository
	ProductRepo        *mock_repositories.MockProductRepository
	StockRepo          *mock_repositories.MockStockRepository
	UploadMetadataRepo *mock_repositories.MockUploadMetadataRepository
}

func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	return &MockRepositories{
		CountryRepo:        mock_repositories.NewMockCountryRepository(ctrl),
		ProductRepo:        mock_repositories.NewMockProductRepository(ctrl),
		StockRepo:          mock_repositories.NewMockStockRepository(ctrl),
		UploadMetadataRepo: mock_repositories.NewMockUploadMetadataRepository(ctrl),
	}

}
