package presentation

import (
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/presentation/handlers"
)

type Handlers struct {
	FileProcessingHandler handlers.FileProcessingHandler
}

func NewHandlers(services *application.Services) *Handlers {
	return &Handlers{
		FileProcessingHandler: *handlers.NewFileProcessingHandler(services),
	}
}
