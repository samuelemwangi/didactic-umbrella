package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadRepository interface {
	UpdateUpload(*domain.FileUploadMetadata) error
}

type uploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) *uploadRepository {
	return &uploadRepository{
		db: db,
	}
}

func (repo *uploadRepository) UpdateUpload(upload *domain.FileUploadMetadata) error {
	result := repo.db.Model(upload).Updates(upload)
	return result.Error
}
