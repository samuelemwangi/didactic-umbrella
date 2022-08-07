package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadRepository interface {
	GetUploadByUploadId(*domain.UploadMetadata) error
	UpdateUploadStatus(*domain.UploadMetadata) error
}

type uploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) *uploadRepository {
	return &uploadRepository{
		db: db,
	}
}

func (repo *uploadRepository) GetUploadByUploadId(upload *domain.UploadMetadata) error {
	result := repo.db.First(upload, "upload_id = ?", upload.UploadID)
	return result.Error

}

func (repo *uploadRepository) UpdateUploadStatus(upload *domain.UploadMetadata) error {
	result := repo.db.Model(&domain.UploadMetadata{}).Where("upload_id = ?", upload.UploadID).Updates(upload)
	return result.Error
}
