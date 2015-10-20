package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/jeffail/gabs"
	"io/ioutil"
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

		body, _ := ioutil.ReadAll(res.Body)
		json, _ := gabs.ParseJSON(body)
		children, _ := json.ChildrenMap()

		for key, child := range children {
			fmt.Printf("Key: %v, Value: %v\n", key, child.Data())
		}
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
