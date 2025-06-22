package services

import (
	"fmt"
	"io"
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


func DeleteObject(bucketName, objectKey string) (int, error) {
	buckets, err := utils.ReadBucketsFromCSV()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	bucketIdx, _ := utils.GetBucketIdx(bucketName, buckets)
	if bucketIdx == -1 {
		return http.StatusNotFound, fmt.Errorf("bucket %s not found", bucketName)
	}

	objectsCSVPath := path.Join(config.Dir, bucketName, "objects.csv")
	objects, err := utils.ReadObjectsFromCSV(objectsCSVPath)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	objectIdx := utils.GetObjectIdx(objectKey, objects)
	if objectIdx == -1 {
		return http.StatusNotFound, fmt.Errorf("object %s not found", objectKey)
	}

	objects.Objects = append(objects.Objects[:objectIdx], objects.Objects[objectIdx+1:]...)
	err = utils.RewriteObjectCSV(objects, objectsCSVPath)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	objectPath := path.Join(config.Dir, bucketName, objectKey)
	err = os.RemoveAll(objectPath)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}