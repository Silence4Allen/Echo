package engine

import (
	"Echo/src/fetcher"
	"fmt"
)

func  Worker(r Request) (ParseResult, error) {
	fmt.Printf("Worker is fetching : %s \r\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		fmt.Printf("Fetcher: Error Fetching Url %s: %v\r\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}
