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
		return nil, utils.ErrBucketAlreadyExists
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

func GetAllBuckets() (*models.Buckets, error) {
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

func DeleteBucket(bucketName string) (int, error) {
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return 505, err
	}

	idx, found := utils.GetBucketIdx(bucketName, buckets)
	if !found {
		return 404, fmt.Errorf("Bucket not found")
	}

	path := path.Join(config.Dir, bucketName)
	err := utils.BucketIsEmtpy(path)

	err = utils.RemoveBucketFromCSV(bucketName)
	if err != nil {
		return 505, err
	}
}
