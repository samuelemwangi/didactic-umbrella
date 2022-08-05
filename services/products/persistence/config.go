package persistence

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type Repositories struct {
	CountryRepo repositories.CountryRepository
	db          *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(dbDriver, dbURL)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &Repositories{
		CountryRepo: repositories.NewCountryRepository(db),
		db:          db,
	}, nil

}

func (r *Repositories) CloseDB() error {
	return r.db.Close()
}

func (r *Repositories) AutoMigrateDB() error {
	return r.db.AutoMigrate(&domain.Country{}).Error
}
