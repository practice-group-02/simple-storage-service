package handlers

import (
	"log"
	"net/http"
	"triple-s/internal/services"
	"triple-s/internal/utils"
)

func CreateObject(w http.ResponseWriter, r *http.Request) {
	op := "PUT /{BucketName}/{ObjectKey}"

	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	object, httpStatus, err := services.CreateObject(bucketName, objectKey, r)
	if err != nil {
		if httpStatus == http.StatusInternalServerError {
			log.Printf("FAIL: OP: %s. Error creating object: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusInternalServerError, "Oops... something went wrong!")
			return
		} else if httpStatus == http.StatusBadRequest || httpStatus == http.StatusNotFound {
			log.Printf("FAIL: OP: %s. Error creating object: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusBadRequest, err.Error())
			return
		} e
	}
}
