package handlers

import (
	"net/http"
	"triple-s/internal/utils"
)

func CraeteObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	object, err := services.CreateObject(bucketName, objectKey)
	if err != nil {
		
	}
})