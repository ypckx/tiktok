package service

import (
	"fmt"
	"tiktok/common"
	"tiktok/model"
)

func FavoriteAction(uid, vid int64, action int8) error {

	//action，1-点赞，2-取消点赞
	if action == 1 {
		err := model.LikeAction(uid, vid)
		if err != nil {
			fmt.Println("LikeAction after error======== :", err.Error())
			return err
		}
	} else {
		err := model.UnLikeAction(uid, vid)
		if err != nil {
			fmt.Println("UnLikeAction after error======== :", err.Error())
			return err
		}
	}
	return nil
}

func FavoriteList(tokenUid, uid int64) (*common.FavoriteListResponse, error) {

	// 获取用户所有点赞视频列表
	favList, err := model.GetFavoriteList(uid)
	if err != nil {
		return nil, err
	}
	// log.Infof("user:%v, followList:%+v", uid, favList)

	favListResponse := common.FavoriteListResponse{
		VideoList: VideoList(favList, tokenUid),
	}

	return &favListResponse, nil
}

func tokenFavList(tokenUserId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := model.GetFavoriteList(tokenUserId)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		m[v.Id] = struct{}{}
	}
	return m, nil
}
