package model

import (
	"errors"
	"fmt"
	"tiktok/config"
	"tiktok/model/db"
	"tiktok/utils"

	"gorm.io/gorm"
)

type Video struct {
	// gorm.Model
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	Author        User   `gorm:"foreignkey:AuthorId"`
}

func (Video) TableName() string {
	return "videos"
}

// 添加用户上传的视频
func InsertVideo(authorid int64, playurl, coverurl, title string) error {
	video := Video{
		AuthorId:      authorid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   utils.GetCurrentTime(),
		Title:         title,
	}
	db := db.GetDB()
	err := db.Create(&video).Error
	if err != nil {
		return err
	}

	// 增加用户的作品数
	var user User
	err = db.Where("user_id = ?", authorid).Find(&user).Error
	if err != nil {
		fmt.Println("update user video count error!")
		return errors.New("update user video count error")
	}

	err = db.Model(&User{}).Where("user_id = ?", authorid).Update("work_count", user.WorkCount+1).Error
	if err != nil {
		fmt.Println("add user video count error")
		return errors.New("add user video count error")
	}

	return nil
}

// 获取视频列表
func GetVideoList(AuthorId int64) ([]Video, error) {
	var videos []Video
	author, err := GetUserInfo(AuthorId)
	if err != nil {
		return videos, err
	}
	db := db.GetDB()
	err = db.Where("author_id = ?", AuthorId).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	for i := range videos {
		videos[i].Author = author
		// 拼接url地址
		videos[i].PlayUrl = config.VideoUrl + videos[i].PlayUrl
		videos[i].CoverUrl = config.CoverUrl + videos[i].CoverUrl
	}
	return videos, nil
}

// 获取视频流，一次最多请求30个
func GetVideoListByFeed(currentTime int64) ([]Video, error) {
	var videos []Video
	db := db.GetDB()
	err := db.Where("publish_time < ?", currentTime).Limit(30).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}

	for i, v := range videos {
		author, err := GetUserInfo(v.AuthorId)


		if err != nil {
			return videos, err
		}
		// 添加video和picture的url
		videos[i].PlayUrl = config.VideoUrl + v.PlayUrl
		videos[i].CoverUrl = config.CoverUrl + v.CoverUrl
		videos[i].Author = author
	}
	return videos, nil
}

