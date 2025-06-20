package services

import (
	"fmt"
	"net/http"
	"triple-s/internal/models"
	"triple-s/internal/utils"
)

func CreateObject(bucketName, objectName string) (*models.Object, int, error) {
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	idx, _ := utils.GetBucketIdx(bucketName, buckets)
	if idx == -1 {
		return nil, http.StatusNotFound, fmt.Errorf("bucket not found from buckets metadata")
	}
}