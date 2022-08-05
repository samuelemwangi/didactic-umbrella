package domain

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Count     int
	CountryID uint
	ProductID uint
	Country   Country
	Product   Product
}
