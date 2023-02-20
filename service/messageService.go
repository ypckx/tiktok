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

	// 获取用户之间发消息的key，通知对方有新信息
	key := common.GetKeyByUserIdAndToUserId(userId, toUserId)
	if _, ok := common.CurMsgInfoMap[key]; !ok {
		common.CurMsgInfoMap[key] = &common.UserMsgInfo{State_HasNewMsg: true, State_NewMsgCount: 0}
	}
	common.CurMsgInfoMap[key].State_NewMsgCount++

	return &common.MessageActionResponse{}, nil
}

func MessageList(userId, toUserId int64, msgTime int64) (*common.MessageListResponse, error) {

	var err error
	var messages []model.Message = nil

	// var key string
	// 获取用户之间发消息的key，用来通知对方是否有新信息
	key := common.GetKeyByUserIdAndToUserId(toUserId, userId)
	if _, ok := common.CurMsgInfoMap[key]; !ok {
		common.CurMsgInfoMap[key] = &common.UserMsgInfo{State_HasNewMsg: false, State_NewMsgCount: 0}
	}

	if msgTime == 0 { // 第一次请求（用户打开聊天界面），获取用户之间的历史聊天记录
		messages, err = model.MessageListCommon(userId, toUserId)
		// key_other := common.GetKeyByUserIdAndToUserId(toUserId, userId)
		// // 别人给我发的消息全部已读
		common.CurMsgInfoMap[key].State_NewMsgCount = 0
	} else {
		// 轮询对方是否给我发新消息
		// key_other := common.GetKeyByUserIdAndToUserId(toUserId, userId)
		msgCount := common.CurMsgInfoMap[key].State_NewMsgCount
		if msgCount == 0 {
			return &common.MessageListResponse{}, nil
		}

		messages, err = model.MessageList(toUserId, userId, msgCount)
		common.CurMsgInfoMap[key].State_NewMsgCount = 0
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
			FromUserId: message.UserId,
			ToUserId:   message.ToUserId,
			Content:    message.Content,
			CreateTime: message.Time,
		}
		list.MessageList[i] = v
		// fmt.Println("userid:", v.FromUserId, "  toUserId:", v.ToUserId, " content:", v.Content)
	}

	return list, nil
}
