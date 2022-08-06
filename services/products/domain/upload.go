package domain

import "github.com/jinzhu/gorm"

const (
	UploadStatusUploaded = iota
	UploadStatusProcessing
	UploadStatusProcessed
)

type FileUploadMetadata struct {
	gorm.Model
	UploadId        string `gorm:"size:50;unique"`
	ProcessedStatus uint
}
