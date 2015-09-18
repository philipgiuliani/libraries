package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	dbUser     = "philipgiuliani"
	dbPassword = ""
	dbName     = "philipgiuliani"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/libraries", LibrariesHandler(db))
	router.GET("/libraries/:libraryID", LibraryHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
