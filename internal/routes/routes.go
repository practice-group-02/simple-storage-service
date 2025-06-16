package routes

import (
	"net/http"
	"triple-s/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /{BucketName}", handlers.CreateBucket)

	return mux
}
