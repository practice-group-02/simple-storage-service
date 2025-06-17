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
			if (name[i] == '-' && name[i-1] == '-') || name[i] == '-' && name[i-1] == '.' || name[i] == '.' && name[i-1] == '-' || name[i] == '.' && name[i-1] == '.'{
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



func ReadBuckets() (*models.Buckets, error){
	file, err := os.OpenFile(path.Join(config.Dir, "buckets.csv"), os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	
	
}