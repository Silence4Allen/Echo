package model

import "Echo/src/util"

type ChannelListWebPage struct {
	Pages    Page      `json:"pages"`
	Channels []Channel `json:"channels"`
}

type ChannelWebPage struct {
	Pages   Page    `json:"pages"`
	Channel Channel `json:"channel"`
	Sounds  []Sound `json:"sounds"`
}

type Page struct {
	PageParam       string   `json:"pageParam"`
	TotalCount      util.Str `json:"totalCount"`
	DefaultPageSize util.Str `json:"defaultPageSize"`
	PageSize        util.Str `json:"_pageSize"`
	PageNum         util.Str `json:"_page"`
}

type Channel struct {
	Id           util.Str `json:"id"`
	Name         string   `json:"name"`
	PicUrl       string   `json:"pic"`
	Info         string   `json:"info"`
	Type         util.Str `json:"type"`
	TagId        util.Str `json:"tag_id"`
	SoundCount   util.Str `json:"sound_count"`
	FollowCount  util.Str `json:"follow_count"`
	LikeCount    util.Str `json:"like_count"`
	ShareCount   util.Str `json:"share_count"`
	UserId       util.Str `json:"user_id"`
	UpdateUserId util.Str `json:"update_user_id"`
	ListOrder    util.Str `json:"list_order"`
	CreateTime   util.Str `json:"create_time"`
	UpdateTime   util.Str `json:"update_time"`
	Status       util.Str `json:"status"`
	Desp         string   `json:"desp"`
	Sound        []Sound  `json:"sound"`
}

type Sound struct {
	Id            util.Str `json:"id"`
	Name          string   `json:"name"`
	Length        util.Str `json:"length"`
	PicUrl        string   `json:"pic"`
	ChannelId     util.Str `json:"channel_id"`
	UserId        util.Str `json:"user_id"`
	Source        string   `json:"source"`
	WebSource     string   `json:"web_source"`
	StatusMask    util.Str `json:"status_mask"`
	CommendTime   util.Str `json:"commend_time"`
	Status        util.Str `json:"status"`
	ShareCount    util.Str `json:"share_count"`
	LikeCount     util.Str `json:"like_count"`
	ExchangeCount util.Str `json:"exchange_count"`
	CommentCount  util.Str `json:"comment_count"`
	ViewCount     util.Str `json:"view_count"`
	IsEdit        util.Str `json:"is_edit"`
	IsPay         util.Str `json:"is_pay"`
	CheckVisition util.Str `json:"check_visition"`
	TranslateMask util.Str `json:"translate_mask"`
	CoverSongId   util.Str `json:"cover_song_id"`
	CoverSongType util.Str `json:"cover_song_type"`
	SoundType     util.Str `json:"sound_type"`
	CreateTime    util.Str `json:"create_time"`
	ParentId      util.Str `json:"parent_id"`
	User          User     `json:"user"`
	Composer      string   `json:"composer"`
	Lyrics        string   `json:"lyrics"`
	OriSinger     string   `json:"ori_singer"`
	SongInfo      SongInfo `json:"song_info"`
	IsBought      util.Str `json:"is_bought"`
}

type User struct {
	Id            util.Str `json:"id"`
	Name          string   `json:"name"`
	Avatar        string   `json:"avatar"`
	Photo         string   `json:"photo"`
	PayClass      util.Str `json:"pay_class"`
	PayStatus     util.Str `json:"pay_status"`
	FamousStatus  util.Str `json:"famous_status"`
	FollowedCount util.Str `json:"followed_count"`
	Status        util.Str `json:"status"`
	IsReady       util.Str `json:"is_ready"`
	TypeMask      util.Str `json:"type_mask"`
	Gender        util.Str `json:"gender"`
	City          string   `json:"city"`
	FamousType    util.Str `json:"famous_type"`
	IsMusician    util.Str `json:"is_musician"`
	IsRealFamous  util.Str `json:"is_real_famous"`
}

type SongInfo struct {
	Name      Info `json:"name"`
	Author    Info `json:"author"`
	AlbumName Info `json:"album_name"`
}

type Info struct {
	Field      string   `json:"field"`
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	VerifyId   util.Str `json:"verify_id"`
	VerifyType util.Str `json:"verify_type"`
}
