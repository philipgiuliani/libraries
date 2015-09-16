package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/libraries", LibrariesHandler).
		Methods("GET")

	router.HandleFunc("/libraries/{libraryId}", LibraryHandler).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
