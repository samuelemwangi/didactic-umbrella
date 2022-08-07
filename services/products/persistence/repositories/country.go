package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryRepository interface {
	SaveCountry(*domain.Country) error
	GetCountries() ([]*domain.Country, error)
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *countryRepository {
	return &countryRepository{
		db: db,
	}
}

func (repo *countryRepository) SaveCountry(country *domain.Country) error {
	result := repo.db.Create(country)
	return result.Error
}

func (repo *countryRepository) GetCountries() ([]*domain.Country, error) {
	var countries []*domain.Country
	result := repo.db.Find(&countries)
	return countries, result.Error
}
