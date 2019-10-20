package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
var nameAndKeyMap = make(map[string]int)


//Country golint wont leave me alone
type Country struct {
	//2-letter ISO format country code
	Code string `json:"alpha2Code"`
	//english human-readable country name
	CountryName string `json:"name"`
	//Country flag
	CountryFlag string `json:"flag"`
	//Species
	Species []string `json:"species"`
	//Species key
	SpeciesKey []int `json:"speciesKey"`
}

//CountryResponse golint wont stop torturing me
type CountryResponse struct {
	Species    string `json:"species"`
	SpeciesKey int    `json:"speciesKey"`
}

//Results should have comment or be unexported
type Results struct {
	Results []CountryResponse `json:"results"`
}

//List a given number of species entries by country

/***Currently it only gets the name, flag and code when the limit is omitted, also picking a lower limit after already getting a result will not omit the previous results*********/
func countryHandler(w http.ResponseWriter, r *http.Request) {
	var itExists = false

	countryIdentifier := r.URL.Path[25:]

	var limit = 20
	if r.URL.Query()["limit"] != nil {
		customLimit := r.URL.Query()["limit"][0]
		customLimitInt, err := strconv.Atoi(customLimit)
		if err == nil {
			limit = customLimitInt
		}
	}
	results := Results{}
	country := Country{}
	//restcountries
	resp, err := http.Get("https://restcountries.eu/rest/v2/alpha/" + countryIdentifier)
	if err != nil {
		//handle error
		fmt.Println("Error parsing request", err)
		return
	}
	defer resp.Body.Close()

	//resp.body JSON
	err = json.NewDecoder(resp.Body).Decode(&country)
	if err != nil {
		//handle error
		fmt.Println("Error reading JSON data", err)
		return
	}
	resp, err = http.Get("http://api.gbif.org/v1/occurrence/search?country=" + countryIdentifier + "&limit=" + strconv.Itoa(limit)) 
	if err != nil {
		//handle error
		fmt.Println("Error parsing request", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&results)
	fmt.Printf("Species key length from json: %d\n", len(results.Results))

	for i := 0; i < len(results.Results); i++ {
		nameAndKeyMap[results.Results[i].Species] = results.Results[i].SpeciesKey
		//Maps can't have duplicate values, silly >:)
	}

	//iterate though the maps, insert the values into to the arrays
	for Species, SpeciesKey := range nameAndKeyMap {
		//check for duplicates in Country with a bool check
		itExists = false
		for i := 0; i < len(country.Species); i++ {
			if country.Species[i] == Species {
				itExists = true
			}
		}
		if !itExists {
			country.Species = append(country.Species, Species)
			country.SpeciesKey = append(country.SpeciesKey, SpeciesKey)
		}

	}
	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(country)
}
