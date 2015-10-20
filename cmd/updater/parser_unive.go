package main

import (
	"net/http"
	"log"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"strconv"
)

type ParserUnive struct {}

func (*ParserUnive) Parse() ([]Library, error) {
	url := "http://static.unive.it/sitows/index/personebiblioteche"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	json, _ := gabs.ParseJSON(body)
	children, _ := json.ChildrenMap()

	libraries := []Library{}
	for key, child := range children {
		takenPlaces, _ := strconv.Atoi(child.Search("persone").Data().(string))
		totalPlaces, _ := strconv.Atoi(child.Search("max").Data().(string))

		library := Library {
			MappingID: key,
			TakenPlaces: takenPlaces,
			TotalPlaces: totalPlaces,
		}
		libraries = append(libraries, library)
	}

	return libraries, nil
}
