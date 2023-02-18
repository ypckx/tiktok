package service

import (
	"fmt"
	"tiktok/common"
	"tiktok/model"
)

func MessageAction(userId, toUserId int64, content, actionType string) (*common.MessageActionResponse, error) {
	// default-only actionType == "1"
	_, err := model.MessageAdd(userId, toUserId, content)

	if err != nil {
		return nil, err
	}

	// 向MsgMap中添加用户信息
	key := common.GetKeyByUserIdAndToUserId(userId, toUserId)
	if v, ok := common.CurMsgInfoMap[key]; ok {
		v.State_NewMsgCount = v.State_NewMsgCount + 1
	} else {
		v = new(common.UserMsgInfo)
		v.State_HasReqFriends = false
		v.State_NewMsgCount = 1
		common.CurMsgInfoMap[key] = v
	}

	return &common.MessageActionResponse{}, nil
}

func MessageList(userId, toUserId int64) (*common.MessageListResponse, error) {

	var err error
	var messages []model.Message = nil

	key := common.GetKeyByUserIdAndToUserId(toUserId, userId)
	if v, ok := common.CurMsgInfoMap[key]; ok {
		if v.State_HasReqFriends {
			messages, err = model.MessageListCommon(userId, toUserId)
			v.State_HasReqFriends = false
		} else {
			if v.State_NewMsgCount == 0 {
				// fmt.Println("v.State_NewMsgCount == 0..........")
				return &common.MessageListResponse{}, nil
			}

			// 和我对话的用户发给我的
			messages, err = model.MessageList(toUserId, userId, v.State_NewMsgCount)
			v.State_NewMsgCount = 0
		}
	} else {
		fmt.Println("[service MessageList] common.CurMsgInfoMap[key] not exist..... userId:", userId, " toUserId:", toUserId)
	}

	if err != nil {
		return nil, err
	}

	list := &common.MessageListResponse{
		MessageList: make([]*common.Message, len(messages)),
	}

	for i, message := range messages {
		v := &common.Message{
			Id:      message.MessageId,
			Content: message.Content,
			// CreateTime: time.Now().Unix(),
			CreateTime: message.Time,
		}
		list.MessageList[i] = v
		// fmt.Println("messageList struct v:", v)
	}

	return list, nil
}
