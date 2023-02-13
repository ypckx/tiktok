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

func FollowAction(userId, toUserId int64) error {
	db := db.GetDB()
	relation := Relation{
		Follow:   userId,
		Follower: toUserId,
	}

	// err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Find(&Relation{}).Error
	result := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Find(&Relation{})

	if result.RowsAffected != 0 {
		return errors.New("you have followed this user")
	}

	err := db.Create(&relation).Error

	if err == nil {
		fmt.Println("关注成功 userid:", userId, "follow touserId:", toUserId, "   OK!")
	} else {
		fmt.Println("关注失败 userid:", userId, "follow touserId:", toUserId, "   Error!")
	}

	var followerUser, followUser User
	err1 := db.Where("user_id = ?", userId).Find(&followUser).Error

	err2 := db.Where("user_id = ?", toUserId).Find(&followerUser).Error

	if err1 != nil || err2 != nil {
		return errors.New("[FollowAction] find user follow error")
	}

	err1 = db.Model(&User{}).Where("user_id = ?", userId).Update("follow_count", followUser.Follow+1).Error
	err2 = db.Model(&User{}).Where("user_id = ?", toUserId).Update("follower_count", followerUser.Follower+1).Error
	if err1 != nil || err2 != nil {
		fmt.Println("[FollowAction] update user follow error")
	}
	// go CacheChangeUserCount(userId, add, "follow")
	// go CacheChangeUserCount(toUserId, add, "follower")
	return nil
}

func UnFollowAction(userId, toUserId int64) error {
	db := db.GetDB()
	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Delete(&Relation{}).Error
	if err != nil {
		fmt.Println("取消关注失败 err:", err.Error())
		return err
	} else {
		fmt.Println("取消关注成功: userId:", userId, "   unfollow:", toUserId)
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

	// log.Debug("unfollow update user cache")
	// go CacheChangeUserCount(userId, sub, "follow")
	// go CacheChangeUserCount(toUserId, sub, "follower")
	return nil
}

func GetFollowList(userId int64, usertype string) ([]User, error) {
	db := db.GetDB()
	re := []Relation{}
	// joinArg := "follower"
	// if usertype == "follower" {
	// 	joinArg = "follow"
	// }
	// err := db.Joins("left join relations on users.user_id = relations."+joinArg+"_id").
	// 	Where("relations."+usertype+"_id = ?", userId).Find(&list).Error
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

// func CacheChangeUserCount(userid, op int64, ftype string) {
// 	uid := strconv.FormatInt(userid, 10)
// 	mutex, _ := common.GetLock("user_" + uid)
// 	defer common.UnLock(mutex)
// 	user, err := CacheGetUser(userid)
// 	if err != nil {
// 		// log.Infof("user:%v miss cache", userid)
// 		return
// 	}
// 	switch ftype {
// 	case "follow":
// 		user.Follow += op
// 	case "follower":
// 		user.Follower += op
// 	case "like":
// 		user.FavCount += op
// 	case "liked":
// 		user.TotalFav += op
// 	}
// 	CacheSetUser(user)
// }
