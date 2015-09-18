package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	// connect to the database
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// routes
	router := httprouter.New()
	router.GET("/libraries", LibrariesHandler(db))
	router.GET("/libraries/:libraryID", LibraryHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
