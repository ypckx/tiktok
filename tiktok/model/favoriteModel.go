package model

import (
	"errors"
	"fmt"
	"tiktok/config"
	"tiktok/model/db"

	"gorm.io/gorm"
)

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func LikeAction(uid, vid int64) error {
	db := db.GetDB()
	favorite := Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	result := db.Where("user_id = ? and video_id = ?", uid, vid).Find(&Favorite{})
	if result.RowsAffected != 0 {
		return errors.New("you have liked this video")
	}

	err := db.Create(&favorite).Error
	if err != nil {
		return err
	}

	// 点赞对视频的影响
	var video Video
	err = db.Where("video_id = ?", vid).Find(&video).Error
	if err != nil {
		return err
	}

	result = db.Model(&Video{}).Where("video_id = ?", vid).Update("favorite_count", video.FavoriteCount+1)
	if result.RowsAffected == 0 {
		return errors.New("update add favorite_count error")
	}

	// fmt.Println("LikeAction .......... favorite uid:", favorite.UserId, " vid:", favorite.VideoId)
	// authorid, _ := CacheGetAuthor(vid)
	// go func() {
	// 	CacheChangeUserCount(uid, add, "like")
	// 	CacheChangeUserCount(authorid, add, "liked")
	// }()
	// go CacheChangeUserCount(uid, add, "like")
	// go CacheChangeUserCount(authorid, add, "liked")
	return nil
}

func UnLikeAction(uid, vid int64) error {
	db := db.GetDB()
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}

	// 取消点赞对视频的影响
	var video Video
	err = db.Where("video_id = ?", vid).Find(&video).Error
	if err != nil {
		return err
	}

	if video.FavoriteCount > 0 {
		result := db.Model(&Video{}).Where("video_id = ?", vid).Update("favorite_count", video.FavoriteCount-1)
		if result.RowsAffected == 0 {
			return errors.New("update sub favorite_count error")
		}
	} else {
		fmt.Println("update sub video favorite_count : video.FavoriteCount == 0 error")
	}

	// fmt.Println("UnLikeAction .......... uid:", uid, " vid:", vid)
	// authorid, _ := CacheGetAuthor(vid)
	// go func() {
	// go CacheChangeUserCount(uid, sub, "like")
	// go CacheChangeUserCount(authorid, sub, "liked")
	// }()
	return nil
}

func GetFavoriteList(uid int64) ([]Video, error) {
	var videos []Video
	db := db.GetDB()

	err := db.Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return []Video{}, nil
	} else if err != nil {
		return nil, err
	}

	for i, v := range videos {
		author, err := GetUserInfo(v.AuthorId)
		if err != nil {
			return videos, err
		}
		videos[i].PlayUrl = config.VideoUrl + videos[i].PlayUrl
		videos[i].CoverUrl = config.CoverUrl + videos[i].CoverUrl
		videos[i].Author = author
		// fmt.Println("++++++++++GetFavoriteList :", videos[i].AuthorId, " ", videos[i].FavoriteCount)
	}
	return videos, nil
}
