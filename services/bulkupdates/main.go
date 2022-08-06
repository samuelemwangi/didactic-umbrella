package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/presentation"
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

	r.GET("/process-file/:fileid", handlers.FileProcessingHandler.ProcessFile)

	// run app
	r.Run(":8086")
}
