package domain

import "time"

type Product struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	SKU       string     `gorm:"size:100;not null;unique" json:"sku"`
	Name      string     `gorm:"size:200;not null;unique" json:"productName"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
