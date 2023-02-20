package common

import "strconv"

// 用户消息时间
type UserMsgInfo struct {
	State_HasNewMsg   bool
	State_NewMsgCount int64
}

var CurMsgInfoMap = map[string]*UserMsgInfo{}

func GetKeyByUserIdAndToUserId(fromUserId, toUserId int64) string {
	return "msg_" + strconv.FormatInt(int64(fromUserId), 10) + "_" + strconv.FormatInt(int64(toUserId), 10)
}
