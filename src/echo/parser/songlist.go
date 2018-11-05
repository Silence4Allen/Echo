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
	sounds := web.Sounds
	result := engine.ParseResult{}
	for i := 0; i < len(sounds); i++ {
		sound := sounds[i]
		id := sound.Id.Str()
		result.Items = append(result.Items, engine.Item{Url: config.SongUrlModel + id, Id: id, Type: sound.Name})
	}


	return result
}
