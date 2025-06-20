package utils

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"triple-s/config"
	"triple-s/internal/models"
)

func ReadRecordsFromCSV(path string) ([][]string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return records, err
	}
	return records, nil
}

func AppendToCSV(filename string, record []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write(record)
}

func ReadBucketsFromCSV() (*models.Buckets, error) {
	file, err := os.OpenFile(path.Join(config.Dir, "buckets.csv"), os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		log.Printf("Empty header of buckets.csv. Need to fix!")
	}
	if len(records) <= 1 {
		return &models.Buckets{Buckets: []models.Bucket{}}, nil
	}

	return RecordsToBuckets(records[1:])
}

func ReadObjectsFromCSV(path string) (*models.Objects, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		log.Printf("Empty header in %s. Need to fix!", path)
	}
	if len(records) <= 1 {
		return &models.Objects{Objects: []models.Object{}}, nil
	}

	return RecordsToObjects(records[1:])
}

func CreateObjectsCSV(path string) error {
	return WriteFileWithHeader(path, objectsHeader)
}
