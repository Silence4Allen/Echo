package main

import (
	"Echo/src/engine"
	"time"
	"fmt"
	"Echo/src/echo/parser"
	"Echo/src/echo/config"
	"Echo/src/echo/scheduler"
	"Echo/src/persist"
)

const url = "http://www.app-echo.com/api/channel/index"

func main() {
	start := time.Now()

	itemSaver, err := persist.ItemSaverByMysql()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
	}

	e.Run(
		engine.Request{
			Url: url,
			ParserFunc: func(contents []byte, url string) engine.ParseResult {
				return parser.ParseFirstPage(contents, url, config.ParseChannelList)
			}})

	last := time.Since(start)
	fmt.Println(last)
}
