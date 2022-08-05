package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryRepository interface {
	SaveCountry(*domain.Country) (*domain.Country, map[string]string)
	GetCountries() ([]domain.Country, map[string]string)
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *countryRepository {
	return &countryRepository{db}
}

func (c *countryRepository) SaveCountry(country *domain.Country) (*domain.Country, map[string]string) {
	errors := make(map[string]string)

	err := c.db.Debug().Create(&country).Error

	if err != nil {
		errors["system_error"] = "an error occured"
		return nil, errors
	}

	return country, nil
}

func (c *countryRepository) GetCountries() ([]domain.Country, map[string]string) {
	errors := make(map[string]string)

	var countries []domain.Country

	err := c.db.Debug().Limit(10).Find(&countries).Error

	if err != nil {
		errors["system_error"] = "an error occured"
	}
	return countries, nil

}
