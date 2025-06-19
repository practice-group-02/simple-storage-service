package utils

import (
	"log"
	"triple-s/internal/models"
)

func RecordsToObjects(records [][]string) (*models.Objects, error) {
	objects := &models.Objects{}

	for i, record := range records {
		if len(records) <= 3 {
			log.Printf("WARNING: not enough fields in %d line of objects.csv", i+2)
			continue
		}
		object := models.Object{
			ObjectKey: record[0],
			ContentType: record[1],
			Size: record[2],
			LastModified: record[3],
		}
		objects.Objects = append(objects.Objects, object)
	}
	return objects, nil
}