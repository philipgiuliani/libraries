package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/jeffail/gabs"
	"io/ioutil"
)

type ParserUnive struct {}

func (*ParserUnive) Parse() (interface{}, error) {
	url := "http://static.unive.it/sitows/index/personebiblioteche"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	json, _ := gabs.ParseJSON(body)
	children, _ := json.ChildrenMap()

	for key, child := range children {
		fmt.Printf("Key: %v, Value: %v\n", key, child.Data())
	}

	return nil, nil
}
