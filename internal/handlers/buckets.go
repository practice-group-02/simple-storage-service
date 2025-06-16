package handlers

import (
	"fmt"
	"net/http"
)

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	fmt.Println(bucketName)
}
