package domain

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	SKU    string  `gorm:"size:100;unique;index"`
	Name   string  `gorm:"size:200;unique"`
	Stocks []Stock 
}
