package main
/*
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Diagnostic golint wont stop hating me
type Diagnostic struct {
	statusGBIF  int
	restCountry int
	version     string
	uptime      int64
}

func diagnosticHandler(w http.ResponseWriter, r *http.Request) {

	var diagnosData Diagnostic

	gbif, err := http.Get("http://api.gbif.org/v1/occurrence/country")
	if err != nil {
		log.Fatalln(err)
	}
	
	euro, err := http.Get("https://restcountries.eu/rest/v2/alpha/VA")
	if err != nil {
		log.Fatalln(err)
	}
	diagnosData.statusGBIF = gbif.StatusCode
	diagnosData.restCountry = euro.StatusCode

	nyTime := time.Now()
	bigTime := nyTime.Sub(upTime)

	fmt.Printf("Diagnostics:\nStatus for gbif: %d\nStatus for restcountry: %d\nVersion: v1\nTime since last restart %d\n",
		diagnosData.statusGBIF, diagnosData.restCountry, bigTime)
}
*/