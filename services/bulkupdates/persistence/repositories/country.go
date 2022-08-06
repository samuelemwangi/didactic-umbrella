package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type CountryRepository interface {
	GetCountry(*domain.Country) error
	SaveCountry(*domain.Country) error
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *countryRepository {
	return &countryRepository{
		db: db,
	}
}

func (repo *countryRepository) GetCountry(country *domain.Country) error {
	result := repo.db.Where(country).Find(country)
	return result.Error
}

func (repo *countryRepository) SaveCountry(country *domain.Country) error {
	result := repo.db.Create(country)
	return result.Error
}
