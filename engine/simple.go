package engine

import (
	"log"
)

type SimpleEngine struct {
}

// Run 运行
func (e SimpleEngine) Run(seeds ...Request) {
	log.Printf("running\n")
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
	}
}
