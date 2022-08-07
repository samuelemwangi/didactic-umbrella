package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadMetadataRepository interface {
	GetUploadByUploadId(*domain.UploadMetadata) error
	UpdateUploadStatus(*domain.UploadMetadata) error
}

type uploadMetadataRepository struct {
	db *gorm.DB
}

func NewUploadMetadataRepository(db *gorm.DB) *uploadMetadataRepository {
	return &uploadMetadataRepository{
		db: db,
	}
}

func (repo *uploadMetadataRepository) GetUploadByUploadId(upload *domain.UploadMetadata) error {
	result := repo.db.First(upload, "upload_id = ?", upload.UploadID)
	return result.Error

}

func (repo *uploadMetadataRepository) UpdateUploadStatus(upload *domain.UploadMetadata) error {
	result := repo.db.Model(&domain.UploadMetadata{}).Where("upload_id = ?", upload.UploadID).Updates(upload)
	return result.Error
}
