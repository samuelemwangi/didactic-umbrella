package domain

import "time"

type Country struct {
	ID        uint16     `json:"id"`
	Name      string     `gorm:"size:100;not null;unique" json:"countryName"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
