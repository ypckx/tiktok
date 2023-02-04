package models

import "time"

/*
为了适配数据库表和响应结构体之间的数据加载

如：VideoView
字段类型为数据中的字段（如果这些字段在数据库表中）
后面的json注释为响应体中的字段

*/
type videoViewResp struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author" gorm:"-"`
	Play_url      string `json:"play_url"`
	Cover_url     string `json:"Cover_url"`
	Like_count    int    `json:"favorite_count"`
	Comment_count int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" gorm:"-"`
	Title         string `json:"title"`
}

type commentViewResp struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"-"`
	VideoId      int64     `json:"-"`
	UserInfo     User      `json:"user" gorm:"-"`
	Content      string    `json:"content"`
	Content_time time.Time `json:"-"`
	Create_date  string    `json:"create_date" gorm:"-"`
}

type userViewResp struct {
	Id             int64  `json:"id"`
	Username       string `json:"name"`
	Follow_count   int64  `json:"follow_count"`
	Follower_count int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow" gorm:"-"`
}

type IdAndToken struct {
	UserId int64
	Token  string
}

// var usersLoginInfo = map[string]User{}
// key-value，存储全部的 id-token
var TokenToIdMap = map[string]int{}

func GenerateToken(username, password string) (userToken string) {
	userToken = username + password
	return userToken
}
