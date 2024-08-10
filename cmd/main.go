package main

import (
	"YouTube-Fetch-API-Backend/internal"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/videos/{page:[0-9]+}/{limit:[0-9]+}", internal.GetVideos).Methods("GET")

	go internal.StartDataFetcher()

	log.Fatal(http.ListenAndServe(":8080", r))
}
