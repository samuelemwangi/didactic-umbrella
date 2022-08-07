package persistence

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	AutoMigrateDB(db)

	return db
}

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&domain.Country{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Stock{})
	db.AutoMigrate(&domain.UploadMetadata{})
	db.Model(&domain.Stock{}).AddForeignKey("country_id", "countries(id)", "RESTRICT", "RESTRICT")
	db.Model(&domain.Stock{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}
