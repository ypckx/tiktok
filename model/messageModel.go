package model

import (
	"tiktok/model/db"
	"time"
)

type Message struct {
	MessageId int64  `gorm:"column:message_id; primary_key;"`
	UserId    int64  `gorm:"column:user_id"`
	ToUserId  int64  `gorm:"column:to_user_id"`
	Content   string `gorm:"column:content"`
	Time      int64  `gorm:"column:time"`
}

func MessageList(userId, toUserId int64, msgCount int64) ([]Message, error) {
	var messages []Message
	db := db.GetDB()
	var err error
	messages = nil

	// 取出最新的消息
	err = db.Where("user_id = ? AND to_user_id = ?", userId, toUserId).Order("time DESC").Limit(int(msgCount)).Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func MessageListCommon(userId, toUserId int64) ([]Message, error) {
	var messages []Message
	db := db.GetDB()
	var err error
	messages = nil

	// 得到userId 和 toUserId 的全部通信消息

	/*
		SELECT * FROM
		(SELECT content,`time` FROM messages
		WHERE (user_id = 3 AND to_user_id = 1) OR (user_id = 1 AND to_user_id = 3)
		) as t ORDER BY `time` LIMIT  2;
	*/

	db.Table("(?) as t", db.Table("messages").
		Select("content,time").
		Where("(user_id = ? AND to_user_id = ?) OR (user_id = ? AND to_user_id = ?)", userId, toUserId, toUserId, userId)).
		Order("time").Find(&messages)

	if err != nil {
		return nil, err
	}

	// fmt.Println("MessageListCommon.........................")
	// for _, v := range messages {
	// 	fmt.Println("all msg ============ content:", v.Content, "  time:", v.Time)
	// }
	return messages, nil
}

func MessageAdd(userId, toUserId int64, content string) (*Message, error) {
	db := db.GetDB()

	// nowtime := time.Now().Format("2006-01-02 15:04:05")
	nowtime := time.Now().Unix()
	message := Message{
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
		Time:     nowtime,
	}
	result := db.Create(&message)
	if result != nil {
		return nil, result.Error
	}

	return &message, nil
}
