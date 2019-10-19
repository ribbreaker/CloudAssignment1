package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Species lint wont leave me alone
type Species struct {
	Key            int32  `json:"key"`
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phylum"`
	Order          string `json:"order"`
	Family         string `json:"family"`
	Genus          string `json:"genusOrAbove"`
	ScientificName string `json:"scientificName"`
	CanonicalName  string `json:"canonicalName"`
	Year           string `json:"year"`
}

//SpeciesYear lint wont leave me alone
type SpeciesYear struct {
	BracketYear string `json:"bracket year"`
	Year        string `json:"year"`
}

func speciesHandler(w http.ResponseWriter, r *http.Request) {

	speciesKey := r.URL.Path[25:]
	println(speciesKey)
	if len(speciesKey) == 0 {
		http.Error(w, "No species key found", http.StatusBadRequest)
	}

	resp, err := http.Get("http://api.gbif.org/v1/species/" + speciesKey)
	if err != nil {
		log.Fatalln(err)
	}

	var speciesData Species
	_ = json.NewDecoder(resp.Body).Decode(&speciesData)

	resp, err = http.Get("http://api.gbif.org/v1/species/" + speciesKey + "/name")
	if err != nil {
		log.Fatalln(err)
	}

	var speciesYear SpeciesYear
	_ = json.NewDecoder(resp.Body).Decode(&speciesYear)

	if speciesYear.BracketYear != "" {
		speciesData.Year = speciesYear.BracketYear
	} else {
		speciesData.Year = speciesYear.Year
	}

	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(speciesData)
}
