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

	"github.com/philipgiuliani/libraries/db"
)

func main() {
	// connect to the database
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	dbconn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()

	// routes
	router := httprouter.New()
	router.GET("/libraries", LibrariesHandler(dbconn))
	router.GET("/libraries/:id", LibraryHandler(dbconn))

	log.Fatal(http.ListenAndServe(":8080", router))
}
