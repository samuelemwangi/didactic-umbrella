package domain

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Quantity  int
	CountryID uint
	ProductID uint
	Country   Country
	Product   Product
}
