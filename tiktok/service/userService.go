package service

import (
	"errors"
	"tiktok/common"
	"tiktok/model"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(userName, password string) (*common.UserRegisterResponse, error) {

	// fmt.Println("[UserRegister] ", userName, "  ", password)
	// return nil, nil
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

	// fmt.Println("token:", token, "  info:=============")

	return registResponse, nil
}

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
	token, err := common.GenToken(info.Id, userName)
	if err != nil {
		return nil, err
	}
	loginResponse := &common.UserLoginResponse{
		UserId: info.Id,
		Token:  token,
	}

	// fmt.Println("[UserLogin++++++++++] user_id:", loginResponse.UserId, " token:", loginResponse.Token)
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

func messageUserInfo(info model.User) *common.User {
	return &common.User{
		Id:             info.Id,
		Name:           info.Name,
		FollowCount:    info.Follow,
		FollowerCount:  info.Follower,
		IsFollow:       false,
		Avatar:         info.Avatar,
		TotalFavorited: info.TotalFav,
		FavoriteCount:  info.FavCount,
	}
}
