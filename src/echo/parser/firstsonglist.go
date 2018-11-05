package parser

import (
	"Echo/src/engine"
	"Echo/src/model"
	"encoding/json"
	"Echo/src/echo/util"
	"Echo/src/util"
)

func ParseFirstSongList(contents []byte, url string) engine.ParseResult {

	//the result of the first song list aim at the song list urls
	web := model.ChannelWebPage{}
	err := json.Unmarshal([]byte(contents), &web)
	if err != nil {
		panic(err)
	}
	result := engine.ParseResult{}
	totalNum := web.Pages.TotalCount.GetStrInt()
	num := web.Pages.PageSize.GetStrInt()
	urls := urlfactory.GetUrlsByCal(url, totalNum, num)
	requests := make([]engine.Request, len(urls))
	for i := 0; i < util.GetPageNum(totalNum, num); i++ {
		request := engine.Request{Url: urls[i], ParserFunc: ParseSongList}
		requests[i] = request
	}
	result.Requests = requests

	return result
}
