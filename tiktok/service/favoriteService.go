package service

import (
	"fmt"
	"tiktok/common"
	"tiktok/model"
)

func FavoriteAction(uid, vid int64, action int8) error {
	if action == 1 {
		// fmt.Println("LikeAction will========")
		// log.Infof("like action uid:%v,vid:%v", uid, vid)
		err := model.LikeAction(uid, vid)
		if err != nil {
			fmt.Println("LikeAction after error======== :", err.Error())
			return err
		}
	} else {
		// fmt.Println("UnLikeAction will========")
		// log.Infof("unlike action uid:%v,vid:%v", uid, vid)
		err := model.UnLikeAction(uid, vid)
		if err != nil {
			fmt.Println("UnLikeAction after error======== :", err.Error())
			return err
		}
	}
	return nil
}

func FavoriteList(tokenUid, uid int64) (*common.FavoriteListResponse, error) {

	// fmt.Println("[module service  func-FavoriteList] tokenUid:", tokenUid, "   uid:", uid)
	// favList, err := model.GetFavoriteList(uid)
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
