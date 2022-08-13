package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type UploadMetadataRepository interface {
	SaveUploadMetaData(uploadMetadata *domain.UploadMetadata) error
}

type uploadMetadataRepository struct {
	db *gorm.DB
}

func NewUploadMetadataRepository(db *gorm.DB) *uploadMetadataRepository {
	return &uploadMetadataRepository{
		db: db,
	}
}

func (repo *uploadMetadataRepository) SaveUploadMetaData(uploadMetadata *domain.UploadMetadata) error {

	result := repo.db.Create(uploadMetadata)
	return result.Error
}
