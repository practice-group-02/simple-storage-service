package services

import (
	"fmt"
	"os"
	"path"
	"time"
	"triple-s/config"
	"triple-s/internal/models"
	"triple-s/internal/utils"
)

func CreateBucket(bucketName string) (*models.Bucket, error) {
	bucketPath := path.Join(config.Dir, bucketName)
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return nil, err
	}

	_, found := utils.GetBucketIdx(bucketName, buckets)
	if found {
		return nil, fmt.Errorf("bucket already exists")
	}

	err = os.Mkdir(bucketPath, 0755)
	if err != nil {
		return nil, err
	}

	newBucket := &models.Bucket{
		Name:         bucketName,
		CreationTime: time.Now().Format(time.RFC3339),
		LastModified: time.Now().Format(time.RFC3339),
		Status:       "ACTIVE",
	}

	err = utils.ParseBucketInMetadata(path.Join(config.Dir, "buckets.csv"), newBucket)

	return newBucket, nil
}
