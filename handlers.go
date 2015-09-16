package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func LibrariesHandler(w http.ResponseWriter, r *http.Request) {
	libraries := Libraries{
		Library{ID: 1, Name: "Bolzano", TakenPlaces: 5, TotalPlaces: 180},
		Library{ID: 2, Name: "Bologna", TakenPlaces: 99, TotalPlaces: 123},
	}

	json.NewEncoder(w).Encode(libraries)
}

func LibraryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	libraryID := vars["libraryId"]

	fmt.Fprintln(w, "Library show:", libraryID)
}
