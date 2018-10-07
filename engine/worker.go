package engine

import (
	"log"

	"../fetcher"
)

func Worker(r Request) (ParseResult, error) {

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, nil
	}

	return r.Parser.Parser(body, r.Url), nil
}
