package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Parser struct {
	URL   string
	Parse func(parser *Parser)
}

var univeParser = &Parser{
	URL:	"http://static.unive.it/sitows/index/personebiblioteche",
	Parse:	func(p *Parser) {
		res, err := http.Get(p.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		var f interface{}
		json.NewDecoder(res.Body).Decode(f)

		fmt.Println(f)
	},
}

var parsers = []*Parser{
	univeParser,
}

func main() {
	for _, p := range parsers {
		fmt.Println("Fetching updates from:", p.URL)
		p.Parse(p)
	}
}
