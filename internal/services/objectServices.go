package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
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

	object := models.Object{
		ObjectKey:    objectName,
		ContentType:  r.Header.Get("Content-Type"),
		Size:         strconv.FormatInt(r.ContentLength, 10),
		LastModified: time.Now().Format(time.RFC3339),
	}
	if object.ContentType == "" {
		object.ContentType = "NoContent"
	}

	objectsCSVPath := path.Join(config.Dir, bucketName, "objects.csv")
	objects, err := utils.ReadObjectsFromCSV(objectsCSVPath)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	objectIdx := utils.GetObjectIdx(objectName, objects)
	log.Printf("HEREEE: %d", objectIdx)
	if objectIdx != -1 {
		err = utils.RewriteExistingObjectCSV(objects, objectIdx, object, objectsCSVPath)
		return nil, http.StatusInternalServerError, err
	} else {
		err = utils.WriteNewObjectInMetaData(&object, objectsCSVPath)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &object, http.StatusOK, nil
}

func GetObjectsOfBucket(bucketName string) (*models.Objects, int, error) {
	if strings.TrimSpace(bucketName) == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("bucket name should not contain only spaces or be empty")
	}

	objectsPath := path.Join(config.Dir, bucketName, "objects.csv")
	objects, err := utils.ReadObjectsFromCSV(objectsPath)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return objects, http.StatusOK, nil
}


func DeleteObject()