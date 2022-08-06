package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
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
	// handlers := presentation.NewHandlers(services)

	// routes
	r := gin.Default()

	// run app
	r.Run(":8085")
}
