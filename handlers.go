package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LibrariesHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query("SELECT id, name, taken_places, total_places FROM libraries")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		libraries := Libraries{}
		for rows.Next() {
			library := Library{}
			err = rows.Scan(&library.ID, &library.Name, &library.TakenPlaces, &library.TotalPlaces)
			if err != nil {
				log.Fatal(err)
			}
			libraries = append(libraries, library)
		}

		json.NewEncoder(w).Encode(libraries)
	}
}

func LibraryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	libraryID := ps.ByName("libraryID")

	fmt.Fprintln(w, "Library show:", libraryID)
}
