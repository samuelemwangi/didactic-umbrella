package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type CountryRepository interface {
	GetCountryByCode(string) (*domain.Country, error)
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

func (repo *countryRepository) GetCountryByCode(countryCode string) (*domain.Country, error) {
	country := &domain.Country{}
	result := repo.db.First(country, "code = ?", country.Code)
	return country, result.Error
}

func (repo *countryRepository) SaveCountry(country *domain.Country) error {
	result := repo.db.Create(country)
	return result.Error
}
