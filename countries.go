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
func countryHandler(w http.ResponseWriter, r *http.Request) {
	countryIdentifier := r.URL.Path[25:]

	// Country codes are two characters long, if not Error
	/*if len(countryIdentifier) != 2 {
		http.Error(w, "Wrong country code used", http.StatusBadRequest)
	}*/

	var limit = 20
	if r.URL.Query()["limit"] != nil {
		customLimit := r.URL.Query()["limit"][0]
		customLimitInt, err := strconv.Atoi(customLimit)
		if err == nil {
			limit = customLimitInt
		}
	}

	//restcountries
	resp, err := http.Get("https://restcountries.eu/rest/v2/alpha/" + countryIdentifier)
	if err != nil {
		//handle error
		fmt.Println("Error parsing request", err)
		return
	}
	defer resp.Body.Close()

	//resp.body JSON
	country := &Country{}
	err = json.NewDecoder(resp.Body).Decode(country)
	if err != nil {
		//handle error
		fmt.Println("Error reading JSON data", err)
		return
	}

	//Loops through all of them and assigns species/specieskey

	//(use map for not getting duplicates)
	for i := 0; i <= limit; i++ {

		resp, err = http.Get(fmt.Sprintf("http://api.gbif.org/v1/occurrence/search?country=%s&limit=%d", countryIdentifier, limit))
		if err != nil {
			//handle error
			fmt.Println("Error parsing request", err)
			return
		}
		defer resp.Body.Close()

		//Secondary struct
		results := &Results{}
		err = json.NewDecoder(resp.Body).Decode(results)
		//countryResponse := CountryResponse{}
		//err = json.NewDecoder(resp.Body).Decode(results)
		//Maps can't have duplicate values, silly >:)

		fmt.Printf("Species length key from json: %d\n", len(results.Results))

		//check for duplicates (maybe with a bool)
		for i := 0; i < len(results.Results); i++ {
			for j := 0; j < len(results.Results); j++ {
				if nameAndKeyMap[results.Results[i].Species] != nameAndKeyMap[results.Results[j].Species] {

				}
			}
			nameAndKeyMap[results.Results[i].Species] = results.Results[i].SpeciesKey

		}

		//country.Species = append(country.Species, results.Results[i].Species)
		//country.SpeciesKey = append(country.SpeciesKey, results.Results[i].SpeciesKey)
		//fmt.Printf("Species key from json: %d\n", len(countryResponse))

		//iterate though the maps, insert the values into to the arrays
		for Species, SpeciesKey := range nameAndKeyMap {
			country.Species = append(country.Species, Species)
			country.SpeciesKey = append(country.SpeciesKey, SpeciesKey)
		}
	}

	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(country)
}

