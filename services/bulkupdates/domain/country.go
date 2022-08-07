package domain

import (
	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	Code   string `gorm:"size:100;unique"`
	Stocks []Stock
}
