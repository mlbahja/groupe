package main

import (
	"fmt"
	"net/http"

	"groupie/Handlers"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", Handlers.IndexHandler)
	http.HandleFunc("/artists/{id}", Handlers.PageHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("http://localhost:8080\nStarting server on :8080")
}
