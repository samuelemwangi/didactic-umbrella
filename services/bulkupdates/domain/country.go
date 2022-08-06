package domain

import (
	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	Name   string `gorm:"size:100;unique"`
	Stocks []Stock
}
