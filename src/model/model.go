package model

import (
	"Echo/src/util"
	"encoding/json"
)

//to calculate the page number
type Page struct {
	PageParam       string   `json:"pageParam"`
	TotalCount      util.Str `json:"totalCount"`
	DefaultPageSize util.Str `json:"defaultPageSize"`
	PageSize        util.Str `json:"_pageSize"`
	PageNum         util.Str `json:"_page"`
}

//to get channel info
type Channel struct {
	Id          util.Str `json:"id"`
	Name        string   `json:"name"`
	PicUrl      string   `json:"pic"`
	Info        string   `json:"info,omitempty"`
	SoundCount  util.Str `json:"sound_count,omitempty"`
	FollowCount util.Str `json:"follow_count,omitempty"`
	LikeCount   util.Str `json:"like_count,omitempty"`
	ShareCount  util.Str `json:"share_count,omitempty"`
}

//to get channel's id
type ChannelListWebPage struct {
	Pages    Page      `json:"pages"`
	Channels []Channel `json:"channels"`
}

//to get sound's id
type ChannelWebPage struct {
	Pages         Page           `json:"pages"`
	Channel       Channel        `json:"channel"`
	ChannelSounds []ChannelSound `json:"sounds"`
}

type ChannelSound struct {
	Id util.Str `json:"id"`
}

type SoundInfoWeb struct {
	SoundInfo Sound `json:"info"`
}

//to get sound info
type Sound struct {
	Id            util.Str `json:"id"`
	Name          string   `json:"name"`
	Length        util.Str `json:"length"`
	PicUrl        string   `json:"pic"`
	Info          string   `json:"info,omitempty"`
	ChannelId     util.Str `json:"channel_id"`
	UserId        util.Str `json:"user_id,omitempty"`
	Source        string   `json:"source"`
	ShareCount    util.Str `json:"share_count,omitempty"`
	LikeCount     util.Str `json:"like_count,omitempty"`
	ExchangeCount util.Str `json:"exchange_count,omitempty"`
	CommentCount  util.Str `json:"comment_count,omitempty"`
	ViewCount     util.Str `json:"view_count,omitempty"`
	IsEdit        util.Str `json:"is_edit,omitempty"`
	IsPay         util.Str `json:"is_pay,omitempty"`
	IsBought      util.Str `json:"is_bought,omitempty"`
}

type SongInfo struct {
	Name      Info `json:"name"`
	Author    Info `json:"author"`
	AlbumName Info `json:"album_name"`
}

type Info struct {
	Field      string   `json:"field,omitempty"`
	Type       string   `json:"type,omitempty"`
	Name       string   `json:"name,omitempty"`
	VerifyId   util.Str `json:"verify_id,omitempty"`
	VerifyType util.Str `json:"verify_type,omitempty"`
}

func GetSoundFromJsonObj(o interface{}) (Sound, error) {
	var sound Sound
	bytes, err := json.Marshal(o)
	if err != nil {
		return sound, err
	}
	err = json.Unmarshal(bytes, &sound)
	return sound, err

}

func GetChnnelFromJsonObj(o interface{}) (Channel, error) {
	var channel Channel
	bytes, err := json.Marshal(o)
	if err != nil {
		return channel, err
	}
	err = json.Unmarshal(bytes, &channel)
	return channel, err

}
