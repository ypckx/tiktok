package common

import "strconv"

type UserMsgInfo struct {
	State_NewMsgCount   int64
	State_HasReqFriends bool
}

// key-value [user_id : UserStateInfo]
// key format: {userId}_{toUserId}
var CurMsgInfoMap = map[string]*UserMsgInfo{}

func GetKeyByUserIdAndToUserId(fromUserId, toUserId int64) string {
	return "msg_" + strconv.FormatInt(int64(fromUserId), 10) + "_" + strconv.FormatInt(int64(toUserId), 10)
}
