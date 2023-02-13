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
