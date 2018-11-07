package config

const (
	ChannelListApiUrlModel   = "http://www.app-echo.com/api/channel/index?page="        //+pageNum
	ChannelApiUrlModelNormal = "http://www.app-echo.com/api/channel/info?id="           //+id
	ChannelApiUrlModelHot    = "http://www.app-echo.com/api/channel/info?order=hot&id=" //+id
	ChannelApiUrlModelNew    = "http://www.app-echo.com/api/channel/info?order=new&id=" //+id
	SongApiUrlModel          = "http://www.app-echo.com/api/sound/info?id="             //+id
	ChnnelUrlModel           = "http://www.app-echo.com/#/channel/"                     //+id
	SongUrlModel             = "http://www.app-echo.com/#/sound/"                       //+id

	ParseChannelList = "ParseChannelList "
	ParseSongList    = "ParseSongList"

	ElasticIndex = "sound"

)
