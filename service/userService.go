package service

import (
	"errors"
	"tiktok/common"
	"tiktok/model"

	"golang.org/x/crypto/bcrypt"
)

// 用户注册
func UserRegister(userName, password string) (*common.UserRegisterResponse, error) {

	err := model.UserNameIsExist(userName)
	if err != nil {
		return nil, err
	}
	info, err := model.InsertUser(userName, password)
	if err != nil {
		return nil, err
	}

	token, err := common.GenToken(info.Id, userName)
	if err != nil {
		return nil, err
	}

	registResponse := &common.UserRegisterResponse{
		UserId: info.Id,
		Token:  token,
	}

	return registResponse, nil
}

// 用户登录
func UserLogin(userName, password string) (*common.UserLoginResponse, error) {
	info, err := model.GetUserInfo(userName)
	if err != nil {
		return nil, err
	}
	//验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password error")
	}

	// 获取用户token
	token, err := common.GenToken(info.Id, userName)
	if err != nil {
		return nil, err
	}
	loginResponse := &common.UserLoginResponse{
		UserId: info.Id,
		Token:  token,
	}

	return loginResponse, nil
}

// 获取登录用户的信息
func UserInfo(userID int64) (*common.UserResponse, error) {
	info, err := model.GetUserInfo(userID)
	if err != nil {
		return nil, err
	}
	user := messageUserInfo(info)

	return &common.UserResponse{User: user}, nil
}

// 将数据库User结构转化为响应User结构
func messageUserInfo(info model.User) *common.User {
	return &common.User{
		Id:              info.Id,
		Name:            info.Name,
		FollowCount:     info.Follow,
		FollowerCount:   info.Follower,
		IsFollow:        false,
		Avatar:          info.Avatar,
		TotalFavorited:  info.TotalFav,
		FavoriteCount:   info.FavCount,
		WorkCount:       info.WorkCount,
		BackGroundImage: info.BackImage,
		Signature:       info.Signature,
	}
}
