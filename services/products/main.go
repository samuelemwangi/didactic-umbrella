package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("loading .env failed")
	}
}

func main() {

	//repos
	repos, err := persistence.NewRepositories()

	// services
	countryService := country.NewCountryService(repos.CountryRepo)
	productService := product.NewProductService(repos.ProductRepo)
	stockService := stock.NewStockService(repos.StockRepo)

	// handlers
	countryHandler := handlers.NewCountryHandler(countryService)
	productHandler := handlers.NewProductHandler(productService)
	stockHandler := handlers.NewStockHandler(stockService)

	if err != nil {
		panic(err)
	}

	defer repos.CloseDB()
	repos.AutoMigrateDB()

	r := gin.Default()

	// Routes
	r.POST("/country", countryHandler.SaveCountry)
	r.GET("/product/:sku", productHandler.GetProductBySKU)
	r.POST("/consume-stock", stockHandler.ConsumeStock)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8888"
	}

	r.Run(":" + app_port)

}
