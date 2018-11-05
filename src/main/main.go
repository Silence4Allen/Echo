package main

import (
	"Echo/src/engine"
	"time"
	"fmt"
	"Echo/src/echo/parser"
	"Echo/src/echo/config"
	"Echo/src/echo/scheduler"
)

const url = "http://www.app-echo.com/api/channel/index"

func main() {
	start := time.Now()
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{}, WorkerCount: 100,
	}

	e.Run(
		engine.Request{url, func(contents []byte, url string) engine.ParseResult {
			return parser.ParseFirstPage(contents, url, config.ParseChannelList)
		}})
	last := time.Since(start)
	fmt.Println(last)
}
