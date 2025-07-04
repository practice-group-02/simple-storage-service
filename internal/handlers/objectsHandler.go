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
		if httpStatus == http.StatusBadRequest || httpStatus == http.StatusNotFound {
			log.Printf("FAIL: OP: %s. Error creating object: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusBadRequest, err.Error())
			return
		} else {
			log.Printf("FAIL: OP: %s. Error creating object: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusInternalServerError, "Oops... something went wrong!")
			return
		}
	}
	log.Printf("OP: %s. Object created successfully!", op)
	utils.WriteXMLResponse(w, http.StatusOK, object)
}

func GetObjectsOfBucket(w http.ResponseWriter, r *http.Request) {
	op := "GET /{BucketName}"

	bucketName := r.PathValue("BucketName")
	objects, httpStatus, err := services.GetObjectsOfBucket(bucketName)
	if err != nil {
		if httpStatus != http.StatusBadRequest {
			log.Printf("FAIL: OP: %s. Error getting objects: %s", op, err)
			utils.ErrXMLResponse(w, httpStatus, err.Error())
			return
		} else {
			log.Printf("FAIL: OP %s. Error getting objects: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusInternalServerError, "Failed to get objects")
			return
		}
	}

	log.Printf("OP: %s. Objects retrieved successfully!", op)
	utils.WriteXMLResponse(w, http.StatusOK, objects)
}


func DeleteObject(w http.ResponseWriter, r *http.Request) {
	op := "DELETE /{BucketName}/{ObjectKey}"

	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	httpStatus, err := services.DeleteObject(bucketName, objectKey)
	if err != nil {
		if httpStatus != http.StatusInternalServerError {
			log.Printf("FAIL: OP: %s. Error deleting object: %s", op, err)
			utils.ErrXMLResponse(w, httpStatus, err.Error())
			return 
		}else {
			log.Printf("FAIL: OP: %s. Error deleting object: %s", op, err)
			utils.ErrXMLResponse(w, http.StatusInternalServerError, "Oops... Something went wrong!")
			return 
		}
	}
	log.Printf("OP: %s. Object %s deleted successfully!", op, objectKey)
	utils.WriteXMLResponse(w, http.StatusNoContent, nil)
}