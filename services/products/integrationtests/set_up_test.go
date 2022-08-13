package integrationtests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../integration-test.env")
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
