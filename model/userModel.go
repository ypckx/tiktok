package model

import (
	"errors"
	"fmt"
	"tiktok/model/db"
	"tiktok/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// gorm.Model
	Id        int64  `gorm:"column:user_id; primary_key;"`
	Name      string `gorm:"column:user_name"`
	Password  string `gorm:"column:password"`
	Follow    int64  `gorm:"column:follow_count"`
	Follower  int64  `gorm:"column:follower_count"`
	Avatar    string `gorm:"column:avatar"`
	BackImage string `gorm:"column:background_image"`
	Signature string `gorm:"column:signatuare"`
	TotalFav  int64  `gorm:"column:total_favorited"`
	WorkCount int64  `gorm:"column:work_count"`
	FavCount  int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

// 检查该用户名是否已经存在
func UserNameIsExist(userName string) error {
	db := db.GetDB()

	user := User{}
	result := db.Where("user_name = ?", userName).Find(&user)

	// err := result.Error
	if (user != User{}) {
		fmt.Println("[UserNameIsExist] username:", userName, "  pwd", user.Password, " name:", user.Name)
	}

	// debug down
	if result.RowsAffected == 0 {
		return nil
	} else {
		fmt.Println("[model - UserNameIsExist] ,username exist")
		return result.Error
	}
}

// 创建用户
func InsertUser(userName, password string) (*User, error) {
	db := db.GetDB()
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	avatarImg := utils.RandomAvatarImg()
	// fmt.Println("InsertUser RandomAvatarImg:", avatarImg)
	user := User{
		Name:      userName,
		Password:  string(hasedPassword),
		Follow:    0,
		Follower:  0,
		TotalFav:  0,
		FavCount:  0,
		WorkCount: 0,
		Avatar:    avatarImg,
		BackImage: utils.RandomBackgroundImg(),
		Signature: utils.RandomSignature(),
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.Error == nil {
		fmt.Println("InsertUser username:", user.Name, " pwd:", user.Password, "   OK!")
	}
	// log.Infof("regist user:%+v", user)
	// go CacheSetUser(user)
	return &user, nil
}

// 获取用户信息
func GetUserInfo(u interface{}) (User, error) {
	db := db.GetDB()
	user := User{}
	var err error
	switch u := u.(type) {
	case int64:
		// user, err = CacheGetUser(u)
		// if err == nil {
		// 	return user, nil
		// }
		err = db.Where("user_id = ?", u).Find(&user).Error

	case string:
		err = db.Where("user_name = ?", u).Find(&user).Error
	default:
		err = errors.New("")
	}
	if err != nil {
		return user, errors.New("user error")
	}
	// go CacheSetUser(user)
	// log.Infof("%+v", user)

	return user, nil
}

// func CacheSetUser(u User) {
// 	uid := strconv.FormatInt(u.Id, 10)
// 	err := common.CacheSet("user_"+uid, u)
// 	if err != nil {
// 		// log.Errorf("set cache error:%+v", err)
// 	}
// }

// func CacheGetUser(uid int64) (User, error) {
// 	key := strconv.FormatInt(uid, 10)
// 	data, err := common.CacheGet("user_" + key)
// 	user := User{}
// 	if err != nil {
// 		return user, err
// 	}
// 	err = json.Unmarshal(data, &user)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
