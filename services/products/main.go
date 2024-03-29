package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/infrastructure/logging"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("loading .env failed")
	}
}

func main() {

	// Db Connection
	db := persistence.OpenDBConnection()
	defer db.Close()

	logger := logging.SetUpLogger()
	defer logger.Sync()

	// wire repositories
	repos := persistence.NewRepositories(db)

	// wire services
	services := application.NewServices(repos)

	// wire handlers
	handlers := presentation.NewHandlers(services)

	// routes
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		v1.GET("/countries", handlers.CountryHandler.GetCountries)

		v1.GET("/products", handlers.ProductHandler.GetProducts)
		v1.GET("/products/:sku", handlers.ProductHandler.GetProductBySKU)

		v1.POST("/stock/consume", handlers.StockHandler.ConsumeStock)

		v1.POST("/upload", handlers.UploadHandler.UploadCSVFile)
	}
	logger.Info("starting server")

	// run app
	r.Run(":8085")
}
