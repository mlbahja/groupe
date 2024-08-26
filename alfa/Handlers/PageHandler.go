package Handlers

import (
	"encoding/json"
	"fmt"
	link "groupie/global"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var ApiData link.ApiOfArtist

func PageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}
	// if r.URL.Path != "/" {
	// 	http.NotFound(w,r)
	// 	return
	// }
	IdNmber, err := strconv.Atoi(r.PathValue("id"))

	if IdNmber == 0 || err != nil  {
		fmt.Fprintf(w, "is not a valid id")
	}
	// here where we get the path that contain artists .
	response1, err := http.Get(link.Api + "/artists/" + r.PathValue("id"))
	// if there is an err with getting file we write it .
	if err != nil {
		log.Printf("Error making GET request of artists: %v", err)
		http.Error(w, "Failed to fetch data from API", http.StatusBadGateway)
		return
	}
	//  this is where awe close the body after open it when we want to get the artists info .
	defer response1.Body.Close()
	var stock link.ArtistData
	// here we transfer the data from json to a datat structer and put on it .
	err = json.NewDecoder(response1.Body).Decode(&stock)

	if err != nil {
		log.Printf("Error decoding JSON of artists: %v", err)
		http.Error(w, "Failed to process data", http.StatusInternalServerError)
		return
	}

	ApiData.ArtistData = stock

	response2, err := http.Get(link.Api + "/locations/" + r.PathValue("id"))

	if err != nil {
		log.Printf("Error making GET request of locations : %v", err)
		http.Error(w, "Failed to fetch data from API", http.StatusBadGateway)
		return
	}

	defer response2.Body.Close()

	var stocklocation link.Locations

	err = json.NewDecoder(response2.Body).Decode(&stocklocation)

	if err != nil {
		log.Printf("Error decoding JSON of locations: %v", err)
		http.Error(w, "Failed to process data", http.StatusInternalServerError)
		return
	}

	ApiData.Locations = stocklocation

	response3, err := http.Get(link.Api + "/dates/" + r.PathValue("id"))
	if err != nil {
		log.Printf("Error making GET request of dates : %v", err)
		http.Error(w, "Failed to fetch data from API", http.StatusBadGateway)
		return
	}

	defer response3.Body.Close()

	var stockDates link.Dates

	err = json.NewDecoder(response3.Body).Decode(&stockDates)

	if err != nil {
		log.Printf("Error decoding JSON of dates: %v", err)
		http.Error(w, "Failed to process data ", http.StatusInternalServerError)
		return
	}

	ApiData.Dates = stockDates

	response4, err := http.Get(link.Api + "/relation/" + r.PathValue("id"))

	if err != nil {
		log.Printf("Error making GET request of relations: %v", err)
		http.Error(w, "Failed to fetch data from API", http.StatusBadGateway)
		return
	}

	defer response4.Body.Close()

	var stockRelation link.Relations
	
	err = json.NewDecoder(response4.Body).Decode(&stockRelation)

	if err != nil {
		log.Printf("Error decoding JSON of relation: %v", err)
		http.Error(w, "Failed to process data", http.StatusInternalServerError)
		return
	}

	ApiData.Relation = stockRelation

	test, err := template.ParseFiles("templates/result.html")
	if err != nil {
		fmt.Fprintf(w, "error in parsing files")
		return
	}

	test.Execute(w, ApiData)

}
