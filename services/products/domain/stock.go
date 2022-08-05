package domain

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Count     int
	CountryID uint16
	ProductID uint64
	Country   Country
	Product   Product
}
