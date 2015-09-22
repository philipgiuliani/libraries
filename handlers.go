package main

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func LibrariesHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		queryParams := r.URL.Query()

		// Parse location data
		var latitude float64
		var longitude float64
		if queryParams.Get("latitude") != "" && queryParams.Get("longitude") != "" {
			latitude, _ = strconv.ParseFloat(queryParams.Get("latitude"), 32)
			longitude, _ = strconv.ParseFloat(queryParams.Get("longitude"), 32)
		}

		var rows *sql.Rows
		var err error
		if latitude != 0 && longitude != 0 {
			rows, err = db.Query(`
					SELECT id, name, taken_places, total_places, earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)) AS distance
					FROM libraries
					ORDER BY distance ASC`,
				latitude, longitude)
		} else {
			rows, err = db.Query("SELECT id, name, taken_places, total_places FROM libraries")
		}

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		libraries := Libraries{}
		for rows.Next() {
			library := Library{}
			err := rows.Scan(&library.ID, &library.Name, &library.TakenPlaces, &library.TotalPlaces, &library.Distance)
			if err != nil {
				log.Fatal(err)
			}
			libraries = append(libraries, library)
		}

		json.NewEncoder(w).Encode(libraries)
	}
}

func LibraryHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		libraryID := ps.ByName("id")

		library := Library{}
		err := db.QueryRow("SELECT id, name, taken_places, total_places, latitude, longitude, description, city, country_code, contact FROM libraries WHERE id = $1", libraryID).
			Scan(&library.ID, &library.Name, &library.TakenPlaces, &library.TotalPlaces, &library.Latitude, &library.Longitude, &library.Description, &library.City, &library.CountryCode, &library.Contact)

		switch {
		case err == sql.ErrNoRows:
			http.NotFound(w, r)
		case err != nil:
			log.Fatal(err)
		default:
			json.NewEncoder(w).Encode(library)
		}
	}
}
