package routes

import (
	"net/http"
	"triple-s/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /{BucketName}", handlers.CreateBucket)
	mux.HandleFunc("GET /", handlers.ListAllBuckets)
	mux.HandleFunc("DELETE /{BucketName}", handlers.DeleteBucket)

	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", handlers.CreateObject)
	mux.HandleFunc("GET /{BucketName}", handlers.GetObjectsOfBucket)
	mux.HandleFunc("DELETE /{BucketName}/{Object}")

	return mux
}
