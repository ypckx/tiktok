package service

import (
	"tiktok/common"
	"tiktok/model"
)

func MessageAction(userId, toUserId int64, content, actionType string) (*common.MessageActionResponse, error) {
	// actionType，1-发送消息
	_, err := model.MessageAdd(userId, toUserId, content)

	if err != nil {
		return nil, err
	}

	return &common.MessageActionResponse{}, nil
}

func MessageList(userId, toUserId int64, msgTime int64) (*common.MessageListResponse, error) {

	var err error
	var messages []model.Message = nil

	// 获取用户之间发消息的key
	key := common.GetKeyByUserIdAndToUserId(toUserId, userId)
	if msgTime == 0 { // 第一次请求（用户打开聊天界面），获取用户之间的历史聊天记录
		// fmt.Println("============= 第一次请求（用户打开聊天界面） MessageList msgTime:", msgTime)
		messages, err = model.MessageListCommon(userId, toUserId)
		common.CurMsgInfoMap[key] = &common.UserMsgInfo{PreMsgTime: msgTime}
	} else {
		// 判断有没有新消息，若有新消息，msgTime不会和上次请求消息时间一致
		// fmt.Println("=============  MessageList msgTime:", msgTime, "  CurMsgInfoMap[key].PreMsgTime:", common.CurMsgInfoMap[key].PreMsgTime)
		if common.CurMsgInfoMap[key].PreMsgTime == 0 || msgTime == common.CurMsgInfoMap[key].PreMsgTime {
			return &common.MessageListResponse{}, nil
		}
		messages, err = model.MessageList(toUserId, userId, msgTime)
		common.CurMsgInfoMap[key].PreMsgTime = msgTime
	}

	if err != nil {
		return nil, err
	}

	// 封装响应数据列表
	list := &common.MessageListResponse{
		MessageList: make([]*common.Message, len(messages)),
	}

	for i, message := range messages {
		v := &common.Message{
			Id:         message.MessageId,
			Content:    message.Content,
			CreateTime: message.Time,
		}
		list.MessageList[i] = v
	}

	return list, nil
}
