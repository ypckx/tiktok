package base_service

import (
	"time"
	"tinyTiktok/models"
)

// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

func FeedService(latest_time string, token string) (retResp models.FeedResp, retErr error) {
	feedListViewResp := models.FeedListDao(models.DB.Statement.Context)
	var videoList []models.Video
	for _, v := range feedListViewResp {
		video := models.Video{
			Id:            v.Id,
			Author:        v.Author,
			PlayUrl:       v.Play_url,
			CoverUrl:      v.Cover_url,
			FavoriteCount: v.Like_count,
			CommentCount:  v.Comment_count,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
		videoList = append(videoList, video)
	}
	retResp.Response = models.Response{StatusCode: 0}
	retResp.NextTime = time.Now().Unix()
	retResp.VideoList = videoList
	return retResp, nil
}
