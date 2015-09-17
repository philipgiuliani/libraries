package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
    router.GET("/libraries", LibrariesHandler)
    router.GET("/libraries/:libraryID", LibraryHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}
