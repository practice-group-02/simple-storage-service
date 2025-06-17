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
	
	err = services.CreateBucket(bucketName)
}
