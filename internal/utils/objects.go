package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"triple-s/internal/models"
)

func GetObjectIdx(objectName string, objects *models.Objects) int {
	for i, object := range objects.Objects {
		if objectName == object.ObjectKey {
			return i
		}
	}
	return -1
}

func RewriteExistingObjectCSV(objects *models.Objects, n int, object *models.Object) {
	objectsCSVPath := path.Join()
	objects.Objects[n] = *object
}

func RecordsToObjects(records [][]string) (*models.Objects, error) {
	objects := &models.Objects{}

	for i, record := range records {
		if len(records) <= 3 {
			log.Printf("WARNING: not enough fields in %d line of objects.csv", i+2)
			continue
		}
		object := models.Object{
			ObjectKey:    record[0],
			ContentType:  record[1],
			Size:         record[2],
			LastModified: record[3],
		}
		objects.Objects = append(objects.Objects, object)
	}
	return objects, nil
}

func ValidateObjectKey(key string) error {
	if len(key) < 1 && len(key) > 1024 {
		return fmt.Errorf("object key must be between 1 and 1024 characters.")
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9\-._*'()]+$`)

	if !re.MatchString(key) {
		return fmt.Errorf("object key contains invalid characters")
	}

	if key[0] == ' ' || key[len(key)-1] == ' ' {
		return fmt.Errorf("object key cannot start or end with spaces")
	}

	for _, ch := range key {
		if ch <= 31 || ch == 127 {
			return fmt.Errorf("object key contains invalid control character: %c", ch)
		}
	}

	invalidChars := []string{"\\", "{", "}", "^", "`", "]", "[", "\"", ">", "<", "#", "|", "%", "~"}
	for _, invalidChar := range invalidChars {
		if contains := stringContains(key, invalidChar); contains {
			return fmt.Errorf("Object key contains invalid character: %s", invalidChar)
		}
	}

	return nil
}

// Helper function to check if a string contains a substring
func stringContains(s, substr string) bool {
	return regexp.MustCompile(`\Q` + substr + `\E`).MatchString(s)
}
