package persist

import (
	"fmt"
	"context"
	"gopkg.in/olivere/elastic.v6"
	"Echo/src/engine"
	"github.com/pkg/errors"
	"database/sql"
	"Echo/src/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
//for elasticsearch to save
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			fmt.Printf("Got Item #%d : Type = %s , Url = %s \r\n", itemCount, item.PayloadType, item.Url)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item saver : error saving item %v:", item)
				continue
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	if err != nil {
		return nil
	}

	return nil

}

const driverName = "mysql"
const dataSourceName = "root:allen@/echo"
const insertStr = "INSERT INTO sound_2(`id`, `name`,`url`,`info`, `length`, `pic_url`, `channel_id`, `user_id`," +
	" `source`, `share_count`, `like_count`, `exchange_count`, `comment_count`, `view_count`, `is_edit`," +
	" `is_pay`, `is_bought`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

func ItemSaverByMysql() (chan engine.Item, error) {
	out := make(chan engine.Item)

	//open the database
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return out, err
	}
	//defer db.Close()

	failedItems := make(map[engine.Item]int)

	go func() {
		itemCount := 0
		for {
			item := <-out
			fmt.Printf("Got Item #%d : Type = %s , Url = %s \r\n", itemCount, item.PayloadType, item.Url)
			itemCount++

			err := insertToMysql(db, item)
			if err != nil {
				log.Printf("Item saver : error saving item which id is %s:", item.Id)
				go func() {
					if failedItems[item] == 0 && failedItems[item] < 10 {
						failedItems[item]++
						out <- item
					}
				}()
				continue
			}
		}
	}()

	return out, nil
}

func insertToMysql(db *sql.DB, item engine.Item) error {
	sound, err := model.GetSoundFromJsonObj(item.Payload)
	if err != nil {
		return err
	}

	_, err = db.Exec(insertStr, sound.Id, sound.Name, item.Url, sound.Info, sound.Length, sound.PicUrl, sound.ChannelId,
		sound.UserId, sound.Source, sound.ShareCount, sound.LikeCount, sound.ExchangeCount,
		sound.CommentCount, sound.ViewCount, sound.IsEdit, sound.IsPay, sound.IsBought)
	if err != nil {
		return err
	}
	fmt.Printf("Mysql: Item which id is %s has been set in the database", item.Id)
	return nil
}
