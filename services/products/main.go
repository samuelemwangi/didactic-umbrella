package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
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
	services := application.NewServices(repos)
	handlers := presentation.NewHandlers(services)

	// routes
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		v1.POST("/country", handlers.CountryHandler.SaveCountry)
		v1.GET("/product/:sku", handlers.ProductHandler.GetProductBySKU)
		v1.POST("/consume-stock", handlers.StockHandler.ConsumeStock)
	}

	// run app
	r.Run(":8085")
}
