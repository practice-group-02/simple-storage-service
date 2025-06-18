package handlers

import (
	"log"
	"net/http"
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
		log.Printf("Fail: OP: %s. Error creating bucket: %s", op, err)
		utils.ErrXMLResponse(w, http.StatusConflict, err.Error())
		return
	}

	log.Printf("OP: %s. Bucket created successfully", op)
	utils.WriteXMLResponse(w, http.StatusOK, bucket)
}
