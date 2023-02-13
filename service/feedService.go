package service

import (
	"tiktok/common"
	"tiktok/model"
	"tiktok/utils"
)

func GetFeedList(currentTime int64, userId int64) (*common.FeedResponse, error) {
	videoList, err := model.GetVideoListByFeed(currentTime)
	if err != nil {
		return nil, err
	}
	feed := &common.FeedResponse{
		VideoList: VideoList(videoList, userId),
	}

	nextTime := utils.GetCurrentTime()
	if len(videoList) == 30 {
		nextTime = videoList[len(videoList)-1].PublishTime
	}
	feed.NextTime = nextTime
	return feed, nil
}

func VideoList(videoList []model.Video, userId int64) []*common.Video {
	var err error
	followList := make(map[int64]struct{})
	favoriteList := make(map[int64]struct{})

	if userId != int64(0) {
		followList, err = tokenFollowList(userId)
		if err != nil {
			return nil
		}
		favoriteList, err = tokenFavList(userId)
		if err != nil {
			return nil
		}
	}

	lists := make([]*common.Video, len(videoList))
	for i, video := range videoList {
		v := &common.Video{
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Author:        messageUserInfo(video.Author),
			Title:         video.Title,
		}
		if _, ok := followList[video.AuthorId]; ok {
			v.Author.IsFollow = true
		}
		if _, ok := favoriteList[video.Id]; ok {
			v.IsFavorite = true
		}
		lists[i] = v
	}

	return lists
}
