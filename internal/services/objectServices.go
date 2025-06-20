package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
	"triple-s/config"
	"triple-s/internal/models"
	"triple-s/internal/utils"
)

func CreateObject(bucketName, objectName string, r *http.Request) (*models.Object, int, error) {
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	idx, _ := utils.GetBucketIdx(bucketName, buckets)
	if idx == -1 {
		return nil, http.StatusNotFound, fmt.Errorf("bucket not found from buckets metadata")
	}

	err = utils.ValidateObjectKey(objectName)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	objectPath := path.Join(config.Dir, bucketName, objectName)
	file, err := os.Create(objectPath)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	object := &models.Object{
		ObjectKey:    objectName,
		ContentType:  r.Header.Get("Content-Type"),
		Size:         strconv.FormatInt(r.ContentLength, 10),
		LastModified: time.Now().Format(time.RFC3339),
	}

	objectsCSVPath := path.Join(config.Dir, bucketName, "objects.csv")
	objects, err := utils.ReadObjectsFromCSV(objectsCSVPath)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	objectIdx := utils.GetObjectIdx(objectName, objects)
	if objectIdx != -1 {
		err = utils.RewriteExistingObjectCSV(objects, objectIdx, object)
	}
}
