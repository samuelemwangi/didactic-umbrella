package domain

import "github.com/jinzhu/gorm"

const (
	UploadStatusUploaded = iota
	UploadStatusProcessing
	UploadStatusProcessingAborted
	UploadStatusProcessed
)

type FileUploadMetadata struct {
	gorm.Model
	FileName           string `gorm:"size:256"`
	UploadId           string `gorm:"size:50;unique"`
	TotalItems         uint
	TotalItemsProcesed uint
	ProcessedStatus    uint
}
