package utils

import (
	"fmt"
)

func ValidateBucketName(name string) error {
	if len(name) < 3 || len(name) > 63 {
		return fmt.Errorf("length of name of bucket should be between 3 and 63")
	}

	for i := 0; i < len(name); i++ {
		if !((name[i] >= 'a' && name[i] <= 'z') || (name[i] >= '0' && name[i] <= '9') || name[i] == '-' || name[i] == '.') {
			return fmt.Errorf("forbidden character in bucket name '%s', character: '%s'", name, string(name[i]))
		}
	}

	return nil
}
