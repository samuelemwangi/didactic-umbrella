package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type UploadRepository interface {
	SaveUploadMetaData(*domain.FileUploadMetadata) (*domain.FileUploadMetadata, *string)
}

type uploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) *uploadRepository {
	return &uploadRepository{
		db: db,
	}
}

func (repo *uploadRepository) SaveUploadMetaData(metadata *domain.FileUploadMetadata) (*domain.FileUploadMetadata, *string) {

	err := repo.db.Debug().Create(&metadata).Error

	if err != nil {
		errorMessage := err.Error()
		return nil, &errorMessage
	}
	return metadata, nil
}
