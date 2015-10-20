package main

import (
	"fmt"
	"log"
)

type Library struct {
	MappingID   string
	TakenPlaces int
	TotalPlaces int
}

type Parser interface {
	Parse() ([]Library, error)
}

func main() {
	parsers := []Parser{
		new(ParserUnive),
	}

	for _, p := range parsers {
		fmt.Println("Fetching updates...")

		libraries, err := p.Parse()
		if err != nil {
			log.Fatal(err)
		}

		for _, library := range libraries {
			fmt.Printf("%v - %v / %v\n", library.MappingID, library.TakenPlaces, library.TotalPlaces)
		}
	}
}
