package utils

import (
	"encoding/xml"
	"log"
	"net/http"
	"triple-s/internal/models"
)

func WriteXMLResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "applicatoin/xml")
	w.WriteHeader(statusCode)

	if err := xml.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding response: %s", err)
	}
}

func ErrXMLResponse(w http.ResponseWriter, code int, message string) {
	WriteXMLResponse(w, code, models.ErrResponse{StatusCode: code, Message: message})
}
