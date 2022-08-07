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

	// do not wait for this to finish
	go handlers.FileProcessingHandler.ProcessFile()

	// routes
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/upload-status/:uploadId", handlers.FileProcessingHandler.GetProcessingStatus)
	}

	// run app
	r.Run(":8086")
}
