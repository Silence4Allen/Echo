package parser

import (
	"Echo/src/engine"
	"Echo/src/model"
	"encoding/json"
	"Echo/src/echo/util"
	"Echo/src/echo/config"
	"Echo/src/util"
)

func ParseFirstPage(contents []byte, url string, name string) engine.ParseResult {

	result := engine.ParseResult{}

	switch name {
	case config.ParseChannelList:
		web := model.ChannelListWebPage{}
		err := json.Unmarshal([]byte(contents), &web)
		if err != nil {
			panic(err)
		}

		totalNum := web.Pages.TotalCount.GetStrInt()
		num := web.Pages.PageSize.GetStrInt()
		urls := urlfactory.GetUrlsByCal(config.ChannelListApiUrlModel, totalNum, num)
		for i := 0; i < util.GetPageNum(totalNum, num); i++ {
			result.Requests = append(result.Requests,
				engine.Request{
					Url:        urls[i],
					ParserFunc: ParseChannelList,
				})
		}

	case config.ParseSongList:
		web := model.ChannelWebPage{}
		err := json.Unmarshal([]byte(contents), &web)
		if err != nil {
			panic(err)
		}

		totalNum := web.Pages.TotalCount.GetStrInt()
		num := web.Pages.PageSize.GetStrInt()
		urls := urlfactory.GetUrlsByCal(url, totalNum, num)
		for i := 0; i < util.GetPageNum(totalNum, num); i++ {
			result.Requests = append(result.Requests,
				engine.Request{
					Url:        urls[i],
					ParserFunc: ParseSongList,
				})
		}
	}

	return result
}
