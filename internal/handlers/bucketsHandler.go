package handlers

import (
	"log"
	"net/http"
	"path"
	"strings"
	"triple-s/config"
	"triple-s/internal/services"
	"triple-s/internal/utils"
)

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	op := "PUT /{BucketName}"
	bucketName := r.PathValue("BucketName")

	err := utils.ValidateBucketName(bucketName)
	if err != nil {
		log.Printf("OP: %s. Validation error: %s", op, err)
		utils.ErrXMLResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bucket, err := services.CreateBucket(bucketName)
	if err != nil {
		if err == utils.ErrBucketAlreadyExists {
			log.Printf("Fail: OP: %s. Error creating bucket: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusConflict, err.Error())
			return
		}
		log.Printf("Fail: OP: %s. Error creating bucket: %s", op, err)
		utils.ErrXMLResponse(w, http.StatusInternalServerError, utils.ErrCreatingBucket.Error())
		return
	}

	err = utils.CreateObjectsCSV(path.Join(config.Dir, bucketName, "objects.csv"))
	if err != nil {
		log.Printf("Fail: OP: %s. Error creating objects.csv", err)
		utils.ErrXMLResponse(w, http.StatusInternalServerError, "failed to create objects metadata")
	}

	log.Printf("OP: %s. Bucket created successfully", op)
	utils.WriteXMLResponse(w, http.StatusOK, bucket)
}

func ListAllBuckets(w http.ResponseWriter, r *http.Request) {
	op := "GET /"
	buckets, err := services.GetAllBuckets()
	if err != nil {
		log.Printf("Fail: OP: %s. Error getting buckets", err)
		utils.ErrXMLResponse(w, http.StatusInternalServerError, "failed to extract buckets")
		return
	}
	log.Printf("OP: %s. Buckets extracted successfully!", op)
	utils.WriteXMLResponse(w, http.StatusOK, buckets)
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if strings.TrimSpace(bucketName) == "" {
		log.Printf("empty bucket name")
		utils.ErrXMLResponse(w, http.StatusBadRequest, "bucket name is empty")
		return 
	}

	bucket, err := services.DeleteBucket(bucketName)
}