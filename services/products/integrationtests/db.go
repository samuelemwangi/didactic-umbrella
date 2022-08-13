package integrationtests

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"

	_ "github.com/go-sql-driver/mysql"
)

func OpenTestDBConnection() *gorm.DB {
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatalln("an error occured while opening db connection", err)
		panic(err)
	}

	db.LogMode(true)

	AutoMigrateTestDB(db)

	return db
}

func AutoMigrateTestDB(db *gorm.DB) {
	db.AutoMigrate(&domain.Country{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Stock{})
	db.AutoMigrate(&domain.UploadMetadata{})
	db.Model(&domain.Stock{}).AddForeignKey("country_id", "countries(id)", "RESTRICT", "RESTRICT")
	db.Model(&domain.Stock{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}
