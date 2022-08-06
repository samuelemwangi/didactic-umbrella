package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadRepository interface {
	UpdateUploadStatus(*domain.FileUploadMetadata) error
}

type uploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) *uploadRepository {
	return &uploadRepository{
		db: db,
	}
}

func (repo *uploadRepository) UpdateUploadStatus(upload *domain.FileUploadMetadata) error {
	result := repo.db.Model(&domain.FileUploadMetadata{}).Where("upload_id = ?", upload.UploadId).Updates(upload)
	return result.Error
}
