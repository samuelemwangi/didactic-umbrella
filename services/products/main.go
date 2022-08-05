package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("loading .env failed")
	}
}

func main() {

	// wire repos
	repos, err := persistence.NewRepositories()

	// Gettting services
	countryService := application.NewCountryService(repos.CountryRepo)

	//Getting handlers
	couuntrHandler := presentation.NewCountryHandler(countryService)

	if err != nil {
		panic(err)
	}

	defer repos.CloseDB()
	repos.AutoMigrateDB()

	r := gin.Default()

	r.POST("/country", couuntrHandler.SaveCountry)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8888"
	}

	r.Run(":" + app_port)

}
