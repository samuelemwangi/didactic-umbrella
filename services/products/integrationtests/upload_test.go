package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/presentation"
	"github.com/segmentio/kafka-go"
)

func TestUploadCSVFile(t *testing.T) {
	db := OpenTestDBConnection()
	defer db.Close()

	prepareEnvironment(db)
	defer cleanEnvironment(db)

	repos := persistence.NewRepositories(db)
	services := application.NewServices(repos)
	handlers := presentation.NewHandlers(services)

	t.Run("UploadCSVFile Test - Upload CSV File", func(t *testing.T) {
		// prepare to upload multipart form
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("file", "test.csv")
		part, _ := writer.CreateFormFile("file", "test.csv")
		part.Write([]byte("country,sku,name,stock_change"))
		part.Write([]byte("\n"))
		part.Write([]byte(`"eg","9befa247cd11","Chung PLC Table","22"`))
		writer.Close()

		gin.SetMode(gin.TestMode)
		r := gin.Default()

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()
		r.POST("/api/v1/upload", handlers.UploadHandler.UploadCSVFile)
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err := json.Unmarshal(rr.Body.Bytes(), &responseMap)

		if err != nil {
			t.Errorf("unable to unmarshal response body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
		}

		consumeMessage, consumeErr := ConsumeMessage("file-process-topic")

		if consumeErr != nil {
			t.Errorf("unable to consume message: %v", consumeErr)
		}

		foundItem := false

		for key, value := range responseMap {
			if reflect.TypeOf(key).Kind() == reflect.String && key == "itemDetails" {
				for _, item := range value.(map[string]interface{}) {
					if item == *consumeMessage {
						foundItem = true
					}
				}

			}
		}

		if !foundItem {
			t.Errorf("expected response to contain %v, got %v", *consumeMessage, responseMap)
		}
	})
}

func prepareEnvironment(db *gorm.DB) {
	// override environment values
	os.Setenv("FILE_PATH", os.Getenv("TEST_FILE_PATH"))
	bootstrapServers := os.Getenv("TEST_BOOTSTRAP_SERVERS")
	if bootstrapServers == "" {
		bootstrapServers = "localhost:29092"
	}
	os.Setenv("BOOTSTRAP_SERVERS", bootstrapServers)
}

func ConsumeMessage(topic string) (*string, error) {

	brokers := []string{os.Getenv("BOOTSTRAP_SERVERS")}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     "kafka-consumer-group-1",
		StartOffset: kafka.FirstOffset,
	})

	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	message := string(m.Value)
	r.Close()
	return &message, nil

}

func cleanEnvironment(db *gorm.DB) {
	filesDirectory := os.Getenv("FILE_PATH")

	// delete all files  that match csv from directory
	files, err := ioutil.ReadDir(filesDirectory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(), ".csv") {
			err := os.Remove(filesDirectory + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}

		}
	}
	db.Unscoped().Delete(&domain.UploadMetadata{})
}
