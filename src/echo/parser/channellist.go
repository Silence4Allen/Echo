package parser

import (
	"Echo/src/model"
	"encoding/json"
	"Echo/src/engine"
	"Echo/src/echo/config"
)

func ParseChannelList(contents []byte, url string) engine.ParseResult {

	//the result of fetch ChannelList
	web := model.ChannelListWebPage{}
	err := json.Unmarshal([]byte(contents), &web)
	if err != nil {
		panic(err)
	}
	//channels
	channels := web.Channels
	urls := make([]string, len(channels))
	result := engine.ParseResult{}
	for i := 0; i < len(channels); i++ {
		channel := channels[i]
		//name := channel.Name
		id := channel.Id
		url := config.ChannelApiUrlModel + id.Str()
		urls[i] = url
		//result.Items = append(result.Items, engine.Item{Type: name, Id: id.Str(), Url: url})
		result.Requests = append(result.Requests, engine.Request{Url: url, ParserFunc: func(contents []byte, url string) engine.ParseResult {
			return ParseFirstPage(contents, url, config.ParseSongList)
		}})
	}

	return result
}
