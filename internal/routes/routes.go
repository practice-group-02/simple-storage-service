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

	return mux
}
