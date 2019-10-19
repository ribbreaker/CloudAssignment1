package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func defaultHandler (w http.ResponseWriter, r *http.Request) {
	helloString := "This is the page of Audun Landøy Solli Bitsec"
	fmt.Printf(helloString)
}


var upTime = time.Now() // this takes a timestamp at the start of the program

func main () {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		//log.Fatal("$PORT must be set")
	}

	//will print port number to heroku logs --tail
	fmt.Println(port)


	//http.HandleFunc("/conservation/v1/", defaultHandler)
	/*http.HandleFunc("/conservation/v1/country/" , countryHandler)
	http.HandleFunc("/conservation/v1/species/", speciesHandler)
	http.HandleFunc("/conservation/v1/diag/", diagHandler)
*/
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
