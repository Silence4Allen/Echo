package persist

import (
	"testing"
	"context"
	"encoding/json"
	"Echo/src/model"
	"Echo/src/engine"
	"gopkg.in/olivere/elastic.v6"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Id:          `1476623`,
		Url:         `http: //www.app-echo.com/#/sound/1476623`,
		Type:        "echo",
		PayloadType: "sound",
		Payload: model.Sound{
			Id:            `1476623`,
			Name:          `桜 「樱花」—Funky治愈系日语`,
			Length:        `320`,
			PicUrl:        `https://qn-up-kibey-echo-cdn.app-echo.com/storage/emulated/0/kibey_echo/photo/image2017_7_8_2_41_8.jpg`,
			ChannelId:     `1115`,
			UserId:        `3462054`,
			Source:        `https://qn-qn-echo-cp-cdn.app-echo.com/c2/98ae889e6aa95b0bf1b6659c22b54536f21f38f6ad67c23282d7a7de4ad8c90c4091efd5.mp3?1502131953`,
			ShareCount:    `0`,
			LikeCount:     `6`,
			ExchangeCount: `0`,
			CommentCount:  `0`,
			ViewCount:     `159`,
			IsEdit:        `2`,
			IsPay:         `0`,
			IsBought:      `1`,
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "sound_test"
	//save item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//read item
	result, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(*result.Source, &actual)

	actualSound, _ := model.GetSoundFromJsonObj(actual.Payload)
	actual.Payload = actualSound

	//compare item
	if actual != expected {
		t.Errorf("Got wrong")
	}
}
