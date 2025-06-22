package services

import (
	"fmt"
	"log"
	"net/http"
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
		return http.StatusInternalServerError, err
	}
	if buckets == nil {
   		return http.StatusInternalServerError, fmt.Errorf("failed to load buckets")
	}

	idx, _ := utils.GetBucketIdx(bucketName, buckets)
	if idx == -1 {
		return http.StatusNotFound, fmt.Errorf("bucket not found")
	}

	objectsPath := path.Join(config.Dir, bucketName, "objects.csv")
	isEmpty, err := utils.BucketIsEmtpy(objectsPath)
	if err != nil && err != utils.ErrBucketIsEmpty {
		return http.StatusInternalServerError, err
	}
	err = nil

	if !isEmpty { 
		return http.StatusConflict, fmt.Errorf("bucket %s is not empty", bucketName)
	}
	if len(buckets.Buckets) <= idx {
		log.Printf("something wrong")
		return http.StatusInternalServerError, fmt.Errorf("problem with buckets length and idx")
	}
	bucket := buckets.Buckets[idx]
	bucketPath := path.Join(config.Dir, bucketName)
	if bucket.Status == "MARKED FOR DELETE" {
		err = os.RemoveAll(bucketPath)
		if err != nil {
			return http.StatusInternalServerError, err
		} 
		err = utils.RemoveBucketFromCSVByIdx(idx)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		err = utils.MarkForDeleteBucketStatus(idx)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	}
	return http.StatusNoContent, nil
}
