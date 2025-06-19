package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"triple-s/config"
	"triple-s/internal/models"
)

func ValidateBucketName(name string) error {
	if len(name) < 3 || len(name) > 63 {
		return fmt.Errorf("length of name of bucket should be between 3 and 63")
	}

	for i := 0; i < len(name); i++ {
		if !((name[i] >= 'a' && name[i] <= 'z') || (name[i] >= '0' && name[i] <= '9') || name[i] == '-' || name[i] == '.') {
			return fmt.Errorf("forbidden character in bucket name '%s', character: '%s'", name, string(name[i]))
		}
		if i != 0 {
			if (name[i] == '-' && name[i-1] == '-') || name[i] == '-' && name[i-1] == '.' || name[i] == '.' && name[i-1] == '-' || name[i] == '.' && name[i-1] == '.' {
				return fmt.Errorf("consecutive hyphens or periods (-- | .. | .- | -.) is not allowed for name of buckets")
			}
		}
	}

	if name[0] < 'a' || name[0] > 'z' {
		return fmt.Errorf("bucket name must start and end with lowercase letter")
	}

	ipPattern := `^(?:\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipPattern, name); matched {
		return fmt.Errorf("bucket name must not be formatted like an IP address")
	}

	prohibitedPrefixes := []string{"xn--", "sthree-", "sthree-configurator", "amzn-s3-demo-"}
	for _, prefix := range prohibitedPrefixes {
		if strings.HasPrefix(name, prefix) {
			return fmt.Errorf("bucket name must not start with the prohibited prefix: " + prefix)
		}
	}

	prohibitedSuffixes := []string{"-s3alias", "--ol-s3", ".mrap", "--x-s3"}
	for _, suffix := range prohibitedSuffixes {
		if strings.HasSuffix(name, suffix) {
			return fmt.Errorf("bucket name must not end with the prohibited suffix: " + suffix)
		}
	}

	return nil
}

func RecordsToBuckets(records [][]string) (*models.Buckets, error) {
	buckets := &models.Buckets{}

	for i, record := range records {
		if len(record) <= 3 {
			return nil, fmt.Errorf("not enough fields in %d line", i+2)
		}
		bucket := models.Bucket{
			Name:         record[0],
			CreationTime: record[1],
			LastModified: record[2],
			Status:       record[3],
		}

		buckets.Buckets = append(buckets.Buckets, bucket)
	}
	return buckets, nil
}

// func BucketsToRecords(buckets *models.Buckets) ([][]string, error) {
// 	records := [][]string{}

// 	for _, bucket := range buckets.Buckets {
// 		record := []string{bucket.Name, bucket.CreationTime, bucket.LastModified, bucket.Status}
// 		records = append(records, record)
// 	}
// }

func GetBucketIdx(bucketName string, buckets *models.Buckets) (int, bool) {
	for i, bucket := range buckets.Buckets {
		if bucket.Name == bucketName {
			return i, true
		}
	}

	return -1, false
}

func ParseBucketInMetadata(metaFilePath string, bucket *models.Bucket) error {
	record := []string{bucket.Name, bucket.CreationTime, bucket.LastModified, bucket.Status}

	err := AppendToCSV(metaFilePath, record)
	if err != nil {
		return err
	}
	return nil
}

func RemoveBucketFromCSVByIdx(idx int) error {
	idx++
	path := path.Join(config.Dir, "buckets.csv")
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file) 
	
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	if _, err = file.Seek(0, 0); err != nil {
        return err
    }
    if err = file.Truncate(0); err != nil {
        return err
    }
	records = append(records[:idx], records[idx+1:]...)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(records)
	
	return nil
}

func BucketIsEmtpy(path string) (bool,error) {
	records, err := ReadRecordsFromCSV(path)
	if err != nil {
		return false, err
	}
	if len(records) > 1 {
		return false, nil
	}
	return true, ErrBucketIsEmpty
}


func MarkForDeleteBucketStatus(idx int) error {
	idx++
	path := path.Join(config.Dir, "buckets.csv")
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	if _, err = file.Seek(0, 0); err != nil {
        return err
    }
    if err = file.Truncate(0); err != nil {
        return err
    }
	records[idx][3] = "MARKED FOR DELETE"
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(records)
	
	return nil
}