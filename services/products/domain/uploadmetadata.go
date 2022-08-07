package domain

import "github.com/jinzhu/gorm"

const (
	UploadStatusUploaded = iota
	UploadStatusProcessing
	UploadStatusProcessingAborted
	UploadStatusProcessed
)

type UploadMetadata struct {
	gorm.Model
	FileName           string `gorm:"size:256"`
	UploadID           string `gorm:"size:50;uniqueIndex"`
	TotalItems         uint
	TotalItemsProcesed uint
	ProcessedStatus    uint
}
