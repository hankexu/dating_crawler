package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s, %v", r.URL, err)
		return ParseResult{}, err
	} else {
		log.Printf("Fetcher: fetching url %s, %v", r.URL, err)
	}

	return r.Parser.Parse(body, r.URL), nil
}
