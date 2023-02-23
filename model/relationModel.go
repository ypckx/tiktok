package model

import (
	"errors"
	"fmt"
	"tiktok/model/db"

	"gorm.io/gorm"
)

const (
	add = int64(1)
	sub = int64(-1)
)

type Relation struct {
	// gorm.Model
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	Follow   int64 `gorm:"column:follow_id"`
	Follower int64 `gorm:"column:follower_id"`
}

func (Relation) TableName() string {
	return "relations"
}

/*
follow_count
关注总数

follower_count
粉丝总数
*/

// 关注操作，自己无法关注自己
func FollowAction(userId, toUserId int64) error {
	db := db.GetDB()
	relation := Relation{
		Follow:   userId,
		Follower: toUserId,
	}

	result := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Find(&Relation{})

	if result.RowsAffected != 0 {
		return errors.New("you have followed this user")
	}

	err := db.Create(&relation).Error

	if err != nil {
		fmt.Println("关注失败 userid:", userId, "follow touserId:", toUserId, "   Error!")
	}

	var followerUser, followUser User
	err1 := db.Where("user_id = ?", userId).Find(&followUser).Error

	err2 := db.Where("user_id = ?", toUserId).Find(&followerUser).Error

	if err1 != nil || err2 != nil {
		return errors.New("[FollowAction] find user follow error")
	}

	// 更新关注用户和被关注用户的数量
	err1 = db.Model(&User{}).Where("user_id = ?", userId).Update("follow_count", followUser.Follow+1).Error
	err2 = db.Model(&User{}).Where("user_id = ?", toUserId).Update("follower_count", followerUser.Follower+1).Error
	if err1 != nil || err2 != nil {
		fmt.Println("[FollowAction] update user follow error")
	}

	return nil
}

// 取消关注
func UnFollowAction(userId, toUserId int64) error {
	db := db.GetDB()
	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Delete(&Relation{}).Error
	if err != nil {
		fmt.Println("取消关注失败 err:", err.Error())
		return err
	}

	var followerUser, followUser User
	err1 := db.Where("user_id = ?", userId).Find(&followUser).Error

	err2 := db.Where("user_id = ?", toUserId).Find(&followerUser).Error

	if err1 != nil || err2 != nil {
		return errors.New("[UnFollowAction] find user follow error")
	}

	if followUser.Follow > 0 {
		err1 = db.Model(&User{}).Where("user_id = ?", userId).Update("follow_count", followUser.Follow-1).Error
	} else {
		fmt.Println("[UnFollowAction] followUser.Follow == 0 when followUser.Follow-1")
	}

	if followerUser.Follower > 0 {
		err2 = db.Model(&User{}).Where("user_id = ?", toUserId).Update("follower_count", followerUser.Follower-1).Error
	} else {
		fmt.Println("[UnFollowAction] followUser.Follow == 0 when followerUser.Follow-1")
	}

	if err1 != nil || err2 != nil {
		fmt.Println("[UnFollowAction] update user follow error")
	}

	return nil
}

func GetFollowList(userId int64, usertype string) ([]User, error) {
	db := db.GetDB()
	re := []Relation{}

	err := db.Where("relations."+usertype+"_id = ?", userId).Find(&re).Error

	if err == gorm.ErrRecordNotFound {
		return []User{}, nil
	} else if err != nil {
		return nil, err
	}
	list := make([]User, len(re))
	for i, r := range re {
		uid := r.Follow
		if usertype == "follow" {
			uid = r.Follower
		}
		list[i], _ = GetUserInfo(uid)
	}
	return list, nil
}

// 获取朋友列表
func GetFriendList(userId int64, usertype string) ([]User, error) {
	/*
		SELECT follow_id FROM
		(
		select
			r1.follow_id,
			r1.follower_id
		FROM
			relations r1  join relations r2
		on r1.follow_id=r2.follower_id  and r1.follower_id=r2.follow_id) as t WHERE follower_id = 1;
	*/
	db := db.GetDB()
	re := []Relation{}

	result := db.Table("(?) as t", db.Table("relations r1").
		Select("r1.follow_id,r1.follower_id").
		Joins("join relations r2 on r1.follow_id = r2.follower_id and r1.follower_id = r2.follow_id")).
		Where("t.follower_id = ?", userId).Find(&re)

	if result.Error != nil {
		fmt.Println("sql err:", result.Error)
	}

	// 获取返回封装的用户列表
	list := make([]User, len(re))
	for i, r := range re {
		v, _ := GetUserInfo(r.Follow)
		list[i] = v
	}

	return list, nil
}
