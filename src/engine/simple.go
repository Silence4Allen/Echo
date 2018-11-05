package engine

import (
	"fmt"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	//the queue of task
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	itemCount := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)

		newRequestsLength := len(parserResult.Requests)
		if newRequestsLength > 0 {
			fmt.Printf("Got %3d new requests , and put them in queue\r\n", newRequestsLength)
		}

		for _, item := range parserResult.Items {
			fmt.Printf("Got Item #%d : name = %s , id = %s\r\n url = %s\r\n", itemCount, item.Type, item.Id, item.Url)
			itemCount++
		}

	}
}
