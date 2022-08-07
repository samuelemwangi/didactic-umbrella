package fileutils

import (
	"encoding/csv"
	"os"
)

type CSVReader interface {
	ReadFile(string) ([][]string, error)
}

type csvReader struct {
}

func NewCSVReader() *csvReader {
	return &csvReader{}
}

func (csvR *csvReader) ReadFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return data, nil

}
