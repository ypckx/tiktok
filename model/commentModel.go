package model

import (
	"fmt"
	"tiktok/model/db"
	"time"
)

type Comment struct {
	// gorm.Model
	CommentId int64  `gorm:"column:comment_id; primary_key;"`
	UserId    int64  `gorm:"column:user_id"`
	VideoId   int64  `gorm:"column:video_id"`
	Comment   string `gorm:"column:comment"`
	Time      string `gorm:"column:time"`
}

func CommentAdd(userId, videoId int64, comment_text string) (*Comment, error) {
	db := db.GetDB()

	nowtime := time.Now().Format("01-02")
	comment := Comment{
		UserId:  userId,
		VideoId: videoId,
		Comment: comment_text,
		Time:    nowtime,
	}
	result := db.Create(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	// 向视频表中添加评论数
	var video Video
	result = db.Where("video_id = ?", videoId).Find(&video)
	if result.RowsAffected == 0 {
		fmt.Println("[Model CommentAdd] videoId not exist in Video")
		return nil, result.Error
	}
	err := db.Model(&Video{}).Where("video_id = ?", videoId).Update("comment_count", video.CommentCount+1).Error
	if err != nil {
		fmt.Println("[Model CommentAdd] update comment_count error")
	}

	// CacheDelCommentAll(videoId)

	return &comment, nil
}

func CommentDelete(videoId, comment_id int64) error {
	db := db.GetDB()
	commentTemp := Comment{}

	err := db.Where("comment_id = ?", comment_id).Take(&commentTemp).Error
	if err != nil {
		return err
	}

	// 在视频中，减少对应评论数
	var video Video
	result := db.Where("video_id = ?", videoId).Find(&video)
	if result.RowsAffected == 0 {
		fmt.Println("[Model CommentDelete] videoId not exist in Video")
		return result.Error
	}

	// 视频评论要存在
	if video.CommentCount > 0 {
		err = db.Model(&Video{}).Where("video_id = ?", videoId).Update("comment_count", video.CommentCount-1).Error
	} else {
		fmt.Println("[Model CommentDelete] video.CommentCount == 0 when sub comment_count")
	}

	if err != nil {
		fmt.Println("[Model CommentDelete] update comment_count error")
	}

	// CacheDelCommentAll(videoId)

	db.Delete(&commentTemp)
	return nil

}

func CommentList(videoId int64) ([]Comment, error) {
	var comments []Comment
	db := db.GetDB()
	var err error

	// comments, _ = CacheGetComment(videoId)
	comments = nil

	// 显示最新评论
	err = db.Where("video_id = ?", videoId).Order("comment_id DESC").Find(&comments).Error

	// CacheSetComment(videoId, comments)
	// log.Infof("comments:%+v", comments)

	if err != nil {
		return nil, err
	}
	return comments, nil

}
