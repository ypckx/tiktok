package models

import (
	"context"
	"time"
	"tinyTiktok/models/db"
)

type VideoView struct {
	Id            int64     `json:"id,omitempty"`
	User_id       int       `json:"user_id"`
	Author        User      `json:"author" gorm:"-"`
	Play_url      string    `json:"play_url,omitempty"`
	Cover_url     string    `json:"cover_url,omitempty"`
	Like_count    int64     `json:"favorite_count,omitempty"`
	Comment_count int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite" gorm:"-"`
	Title         string    `json:"title,omitempty"`
	Create_time   time.Time `json:"create_time"`
}

func FeedListDao(ctx context.Context) []VideoView {
	var videoList []VideoView

	// 查询所有视频
	conn := DB.WithContext(ctx).Model(&db.Video{}).Find(&videoList)

	// 查询视频的作者信息
	for index, _ := range videoList {
		var user User
		conn.Model(&db.User{}).Where("id = ?", videoList[index].User_id).Find(&user)
		videoList[index].Author = user
	}
	return videoList
}
