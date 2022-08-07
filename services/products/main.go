package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/infrastructure/queueing"
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

	// wire repositories
	repos := persistence.NewRepositories(db)
	services := application.NewServices(repos, queueing.NewKafkaProducer())
	handlers := presentation.NewHandlers(services)

	// routes
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		v1.POST("/country", handlers.CountryHandler.SaveCountry)
		v1.GET("/product/:sku", handlers.ProductHandler.GetProductBySKU)
		v1.POST("/consume-stock", handlers.StockHandler.ConsumeStock)
		v1.POST("/upload", handlers.UploadHandler.UploadCSVFile)
	}

	// run app
	r.Run(":8085")
}
