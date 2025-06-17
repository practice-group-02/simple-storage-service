package services

import (
	"os"
	"path"
	"triple-s/config"
	"triple-s/internal/utils"
)

func CreateBucket(bucketPath string) error {
	file, err := os.OpenFile(path.Join(config.Dir, "buckets.csv"), os.O_RDWR, 0755) 
	if err != nil {
		return err
	}
	defer file.Close()

	buckets, err := utils.ReadBuckets()

	err = os.Mkdir(bucketPath, 0755)
	if err != nil {
		return err
	}
}

