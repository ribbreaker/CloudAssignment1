package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Diagnostic golint wont stop hating me
type Diagnostic struct {
	StatusGBIF  int    `json:"gbif"`
	RestCountry int    `json:"restcountries"`
	Version     string `json:"version"`
	Uptime      int    `json:"uptime"`
}

var upTime = time.Now()

func diagnosticHandler(w http.ResponseWriter, r *http.Request) {

	gbif, err := http.Get("http://api.gbif.org/v1/")
	if err != nil {
		log.Fatalln(err)
	}

	euro, err := http.Get("https://restcountries.eu/rest/v2/")
	if err != nil {
		log.Fatalln(err)
	}

	gbifStatus := gbif.StatusCode
	restCountryStatus := euro.StatusCode

	nyTime := time.Now()
	bigTime := int(nyTime.Sub(upTime) / time.Second)

	var diagnosData = Diagnostic{gbifStatus, restCountryStatus, "v1", bigTime}

	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(diagnosData)
}
