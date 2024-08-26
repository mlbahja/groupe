package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	link "groupie/global"
)

// this function is for the first web page , it's desiplay all the data of artists in one page

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// this the variable that we stor on it  all the artists informations
	var DATA []link.ArtistData

	response, err := http.Get(link.Api + "/artists")
	if err != nil {
		fmt.Fprintf(w, "an error of geting link")
		return
	}

	// this is make sure that we close the body , like if there is an open ressourse , it make sure they colsed, like connection of netwerk.
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&DATA)
	if err != nil {
		fmt.Fprintf(w, "an error in decode data ")
		return
	}

	test, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "error in parsing files")
		return
	}

	if err := test.Execute(w, DATA); err != nil {
		fmt.Fprintf(w, "error executing files")
	}
}
