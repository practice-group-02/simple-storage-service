package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"triple-s/config"
	"triple-s/internal/routes"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatal("Error parsing flags: ", err)
	}

	err = config.InitDir()
	if err != nil {
		log.Fatal("Initializing directory error: ", err)
	}

	router := routes.NewRouter()

	addr := ":" + strconv.Itoa(config.Port)
	fmt.Printf("Running server on http://localhost%s\n", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
