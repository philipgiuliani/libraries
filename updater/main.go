package main

import (
	"fmt"
	"log"
)

type Parser interface {
	Parse() (interface{}, error)
}

var parsers = []Parser{
	new(ParserUnive),
}

func main() {
	for _, p := range parsers {
		fmt.Println("Fetching updates...")
		if result, err := p.Parse(); err != nil {
			fmt.Println(result)
		} else {
			log.Fatal(err)
		}
	}
}
