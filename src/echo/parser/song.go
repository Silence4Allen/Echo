package parser

import (
	"Echo/src/engine"
	"Echo/src/model"
	"encoding/json"
	"Echo/src/echo/config"
)

func ParseSong(contents []byte, url string) engine.ParseResult {
	web := model.SoundInfoWeb{}
	err := json.Unmarshal(contents, &web)
	if err != nil {
		panic(err)
	}

	songInfo := web.SoundInfo
	result := engine.ParseResult{}
	result.Items = append(result.Items,
		engine.Item{
			//Name:           songInfo.Name,
			Id:          songInfo.Id.Str(),
			Url:         config.SongUrlModel + songInfo.Id.Str(),
			Type:        "echo",
			PayloadType: "sound",
			//ItemSourcePath: url,
			Payload: songInfo,
		})

	return result
}
