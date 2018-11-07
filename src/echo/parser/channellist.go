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
	result := engine.ParseResult{}
	for i := 0; i < len(channels); i++ {
		channel := channels[i]
		channelID := channel.Id.Str()

		urlNormal := config.ChannelApiUrlModelNormal + channelID
		urlHot := config.ChannelApiUrlModelHot + channelID
		urlNew := config.ChannelApiUrlModelNew + channelID

		result = appendRequests(urlNormal, result)
		result = appendRequests(urlHot, result)
		result = appendRequests(urlNew, result)
		//channel info
		//result.Items = append(result.Items,
		//	engine.Item{
		//		//Name:           channel.Name,
		//		Id:          channelID,
		//		Url:         config.ChnnelUrlModel + channelID,
		//		Type:        "echo",
		//		PayloadType: "channel",
		//		//ItemSourcePath: url,
		//		Payload: channel,
		//	})
	}

	return result
}

func appendRequests(url string, result engine.ParseResult) engine.ParseResult {
	result.Requests = append(result.Requests,
		engine.Request{
			Url: url,
			ParserFunc: func(contents []byte, url string) engine.ParseResult {
				return ParseFirstPage(contents, url, config.ParseSongList)
			}})
	return result
}
