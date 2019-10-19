package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

//Indicated the availability of individual services this service depends on
func diagnosticHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Diagnostic Handler, %q", html.EscapeString(r.URL.Path))
	resp, err := http.Get(r.RequestURI)
	if err != nil {
		//handle error
	}
	resp.Body.Close()
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port ="8080"
	}

	//http.HandleFunc("/conservation/v1/species/", speciesHandler)
	http.HandleFunc("/conservation/v1/country/", countryHandler)
	//	http.HandleFunc("/conservation/v1/diag/", diagnosticHandler)

	//for testing
	//log.Fatal(http.ListenAndServe(":8080", nil))
	//for heroku
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
