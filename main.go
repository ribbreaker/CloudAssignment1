package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/conservation/v1/species/", speciesHandler)
	http.HandleFunc("/conservation/v1/country/", countryHandler)
	http.HandleFunc("/conservation/v1/diag/", diagnosticHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
