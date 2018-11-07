package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"Echo/src/model"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"Echo/src/engine"
)

func main() {
	item := engine.Item{
		Id:          `1476623`,
		Url:         `http://www.app-echo.com/#/sound/1476623`,
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
	failedItems := make(map[engine.Item]int, 10)
	if failedItems[item] == 0 && failedItems[item] < 5 {
		fmt.Println(failedItems[item])
		failedItems[item] = failedItems[item] + 1
	} else {
		fmt.Println(failedItems[item])
	}
	//in := ItemSaverByMysql()
	//time.Sleep(5 * time.Second)
	//in <- engine.Item{
	//	Id:          `1476623`,
	//	Url:         `http://www.app-echo.com/#/sound/1476623`,
	//	Type:        "echo",
	//	PayloadType: "sound",
	//	Payload: model.Sound{
	//		Id:            `1476623`,
	//		Name:          `桜 「樱花」—Funky治愈系日语`,
	//		Length:        `320`,
	//		PicUrl:        `https://qn-up-kibey-echo-cdn.app-echo.com/storage/emulated/0/kibey_echo/photo/image2017_7_8_2_41_8.jpg`,
	//		ChannelId:     `1115`,
	//		UserId:        `3462054`,
	//		Source:        `https://qn-qn-echo-cp-cdn.app-echo.com/c2/98ae889e6aa95b0bf1b6659c22b54536f21f38f6ad67c23282d7a7de4ad8c90c4091efd5.mp3?1502131953`,
	//		ShareCount:    `0`,
	//		LikeCount:     `6`,
	//		ExchangeCount: `0`,
	//		CommentCount:  `0`,
	//		ViewCount:     `159`,
	//		IsEdit:        `2`,
	//		IsPay:         `0`,
	//		IsBought:      `1`,
	//	},
	//}

}

const driverName = "mysql"
const dataSourceName = "root:allen@/echo"
const queryStr = "SELECT * FROM sound;"
const insertStr = "INSERT INTO sound(`id`, `name`,`url`,`info`, `length`, `pic_url`, `channel_id`," +
	" `user_id`, `source`, `share_count`, `like_count`, `exchange_count`, `comment_count`, `view_count`, `is_edit`," +
	" `is_pay`, `is_bought`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

var Id, Name, Url, Info, Length, PicUrl, ChannelId, UserId, Source, ShareCount, LikeCount, ExchangeCount, CommentCount, ViewCount, IsEdit, IsPay, IsBought string

func ItemSaverByMysql() chan engine.Item {
	in := make(chan engine.Item)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	go func() {
		for {
			fmt.Println("routine")
			item := <-in
			fmt.Printf("Got Item : Type = %s , Url = %s \r\n", item.PayloadType, item.Url)
			insertData(db, item)
			query(db)
		}
	}()
	return in
}

func insertData(db *sql.DB, item engine.Item) {
	sound, err := model.GetSoundFromJsonObj(item.Payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(sound)
	tx, _ := db.Begin()

	stmt, err := tx.Prepare(insertStr)
	defer stmt.Close()

	fmt.Printf("Id=%s, Name=%s, Url=%s,Info=%s,Length=%s, PicUr=%s,  ChannelId=%s, UserId=%s, Source=%s, ShareCount=%s,"+
		" LikeCount=%s, ExchangeCount=%s, CommentCount=%s, ViewCount=%s, IsEdit=%s,IsPay=%s, IsBought=%s\r\n", sound.Id, sound.Name, item.Url, sound.Info, sound.Length, sound.PicUrl, sound.ChannelId,
		sound.UserId, sound.Source, sound.ShareCount, sound.LikeCount, sound.ExchangeCount,
		sound.CommentCount, sound.ViewCount, sound.IsEdit, sound.IsPay, sound.IsBought)

	_, err = stmt.Exec(sound.Id, sound.Name, item.Url, sound.Info, sound.Length, sound.PicUrl, sound.ChannelId,
		sound.UserId, sound.Source, sound.ShareCount, sound.LikeCount, sound.ExchangeCount,
		sound.CommentCount, sound.ViewCount, sound.IsEdit, sound.IsPay, sound.IsBought)

	if err != nil {
		log.Error("insert data to database got wrong , item = %v", sound)
		//panic(err)
	}

	tx.Commit()
}

func query(db *sql.DB) {
	db.QueryRow(queryStr).Scan(&Id, &Name, &Url, &Info, &Length, &PicUrl, &ChannelId, &UserId, &Source, &ShareCount,
		&LikeCount, &ExchangeCount, &CommentCount, &ViewCount, &IsEdit, &IsPay, &IsBought)

	fmt.Printf("Id=%s, Name=%s, Url=%s,Info=%s,Length=%s, PicUr=%s,  ChannelId=%s, UserId=%s, Source=%s, ShareCount=%s,"+
		" LikeCount=%s, ExchangeCount=%s, CommentCount=%s, ViewCount=%s, IsEdit=%s,IsPay=%s, IsBought=%s\r\n",
		Id, Name, Url, Info, Length, PicUrl, ChannelId, UserId, Source, ShareCount,
		LikeCount, ExchangeCount, CommentCount, ViewCount, IsEdit, IsPay, IsBought)

}
