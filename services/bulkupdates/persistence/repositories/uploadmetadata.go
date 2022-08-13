package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type UploadMetadataRepository interface {
	GetUploadByUploadId(string) (*domain.UploadMetadata, error)
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

func (repo *uploadMetadataRepository) GetUploadByUploadId(uploadID string) (*domain.UploadMetadata, error) {
	upload := &domain.UploadMetadata{}
	result := repo.db.First(upload, "upload_id = ?", uploadID)
	return upload, result.Error

}

func (repo *uploadMetadataRepository) UpdateUploadStatus(upload *domain.UploadMetadata) error {
	itemsToUpdate := map[string]interface{}{
		"processed_status":     upload.ProcessedStatus,
		"total_items":          upload.TotalItems,
		"total_items_procesed": upload.TotalItemsProcesed,
	}

	result := repo.db.Model(&domain.UploadMetadata{}).Where("upload_id = ?", upload.UploadID).Updates(itemsToUpdate)
	return result.Error
}
