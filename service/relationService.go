package service

import (
	"errors"
	"fmt"
	"tiktok/common"
	"tiktok/model"
)

// 用户关注操作
func RelationAction(toUserId, tokenUserId int64, action string) error {
	// 自己无法关注自己
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

// 获取我关注的列表
func RelationFollowList(userId int64, tokenUserId int64) (*common.FollowListResponse, error) {
	followList, err := model.GetFollowList(userId, "follow")
	if err != nil {
		fmt.Println("service-RelationFollowList model.GetFollowList error!")
		return nil, err
	}

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

// 获取关注我（粉丝）的用户列表
func RelationFollowerList(userId int64, tokenUserId int64) (*common.FollowerListResponse, error) {
	followList, err := model.GetFollowList(userId, "follower")
	if err != nil {
		fmt.Println("service-RelationFollowerList model.GetFollowList error!")
		return nil, err
	}

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

// 获取朋友列表（用户要相互关注才会成为朋友）
func RelationFriendList(userId int64, tokenUserId int64) (*common.FriendListResponse, error) {
	// 获取用户朋友列表
	friendList, err := model.GetFriendList(userId, "follower")
	if err != nil {
		fmt.Println("service-RelationFollowerList model.GetFollowList error!")
		return nil, err
	}

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
		// key := common.GetKeyByUserIdAndToUserId(u.Id, userId)
		// if item, ok := common.CurMsgInfoMap[key]; !ok {
		// 	v := new(common.UserMsgInfo)
		// 	v.State_HasReqFriends = true
		// 	v.State_NewMsgCount = 0
		// 	common.CurMsgInfoMap[key] = v
		// } else {
		// 	item.State_HasReqFriends = true
		// }
	}

	return &friendListResponse, nil
}
