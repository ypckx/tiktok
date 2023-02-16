package service

import (
	"errors"
	"fmt"
	"tiktok/common"
	"tiktok/model"
)

func RelationAction(toUserId, tokenUserId int64, action string) error {
	if tokenUserId == toUserId {
		return errors.New("you can't follow yourself")
	}
	if action == "1" {
		// log.Infof("follow action id:%v,toid:%v", tokenUserId, toUserId)
		err := model.FollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	} else {
		// log.Infof("unfollow action id:%v,toid:%v", tokenUserId, toUserId)
		err := model.UnFollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

func RelationFollowList(userId int64, tokenUserId int64) (*common.FollowListResponse, error) {
	followList, err := model.GetFollowList(userId, "follow")
	if err != nil {
		fmt.Println("service-RelationFollowList model.GetFollowList error!")
		return nil, err
	}
	// log.Infof("user:%v, followList:%+v", userId, followList)
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	followListResponse := common.FollowListResponse{
		UserList: make([]*common.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func RelationFollowerList(userId int64, tokenUserId int64) (*common.FollowerListResponse, error) {
	followList, err := model.GetFollowList(userId, "follower")
	if err != nil {
		fmt.Println("service-RelationFollowerList model.GetFollowList error!")
		return nil, err
	}
	// log.Infof("user:%v, followerList:%+v", userId, followList)
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	followListResponse := common.FollowerListResponse{
		UserList: make([]*common.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func tokenFollowList(userId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := model.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		m[u.Id] = struct{}{}
	}
	return m, nil
}

// 朋友列表要相互关注
func RelationFriendList(userId int64, tokenUserId int64) (*common.FriendListResponse, error) {
	friendList, err := model.GetFriendList(userId, "follower")
	if err != nil {
		fmt.Println("service-RelationFollowerList model.GetFollowList error!")
		return nil, err
	}

	// log.Infof("user:%v, followerList:%+v", userId, followList)
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	friendListResponse := common.FriendListResponse{
		UserList: make([]*common.User, len(friendList)),
	}
	for i, u := range friendList {
		follow := messageUserInfo(u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		friendListResponse.UserList[i] = follow

		// 设置 userId 可以得到全部消息
		key := common.GetKeyByUserIdAndToUserId(u.Id, userId)
		if item, ok := common.CurMsgInfoMap[key]; !ok {
			v := new(common.UserMsgInfo)
			v.State_HasReqFriends = true
			v.State_NewMsgCount = 0
			common.CurMsgInfoMap[key] = v
		} else {
			item.State_HasReqFriends = true
		}
	}

	return &friendListResponse, nil
}
