package parser

import (
	"Echo/src/model"
	"encoding/json"
	"Echo/src/engine"
	"Echo/src/echo/config"
	"Echo/src/util"
	"Echo/src/echo/util"
)

func ParseHomePage(contents []byte, url string) engine.ParseResult {

	//the result of fetch FirstPage aim at the channel list urls
	web := model.ChannelListWebPage{}
	err := json.Unmarshal([]byte(contents), &web)
	if err != nil {
		panic(err)
	}

	result := engine.ParseResult{}
	totalNum := web.Pages.TotalCount.GetStrInt()
	num := web.Pages.PageSize.GetStrInt()
	urls := urlfactory.GetUrlsByCal(config.ChannelListApiUrlModel, totalNum, num)
	requests := make([]engine.Request, len(urls))
	for i := 0; i < util.GetPageNum(totalNum, num); i++ {
		request := engine.Request{Url: urls[i], ParserFunc: ParseChannelList}
		requests[i] = request
	}
	result.Requests = requests
	//dom't need this result Item data
	return result
}
