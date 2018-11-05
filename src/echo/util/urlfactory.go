package urlfactory

import (
	"strconv"
	"Echo/src/echo/config"
	"Echo/src/util"
)

//dealwith urls of the channel list and the song list
func GetUrlsByCal(url string, totalNum int, num int) []string {
	pageNum := util.GetPageNum(totalNum, num)
	var urls []string
	if url == config.ChannelListApiUrlModel { //channel list
		for i := 0; i < pageNum; i++ {
			url := url + strconv.Itoa(i+1)
			urls = append(urls, url)
		}
	} else { //song list
		for i := 0; i < pageNum; i++ {
			url := url + "&page=" + strconv.Itoa(i+1)
			urls = append(urls, url)
		}
	}

	return urls
}
