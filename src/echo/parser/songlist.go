package parser

import (
	"Echo/src/engine"
	"Echo/src/model"
	"encoding/json"
	"Echo/src/echo/config"
)

func ParseSongList(contents []byte, url string) engine.ParseResult {

	web := model.ChannelWebPage{}
	err := json.Unmarshal([]byte(contents), &web)
	if err != nil {
		panic(err)
	}

	//songs
	channelSounds := web.ChannelSounds
	result := engine.ParseResult{}
	for i := 0; i < len(channelSounds); i++ {
		channelSound := channelSounds[i]
		id := channelSound.Id.Str()
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        config.SongApiUrlModel + id,
				ParserFunc: ParseSong,
			})
	}

	return result
}
