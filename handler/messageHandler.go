package handler

import (
	"fmt"
	"strconv"
	"tiktok/response"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func MessageAction(ctx *gin.Context) {

	var err error
	tokenUids, _ := ctx.Get("UserId")

	tokenUid := tokenUids.(int64)

	to_user_id := ctx.Query("to_user_id")
	content := ctx.Query("content")
	actionType := ctx.Query("action_type")

	// 发送消息
	toUserId, err := strconv.ParseInt(to_user_id, 10, 64)
	if err != nil {
		fmt.Println("handler func-MessageAction toUserId error:", err.Error())
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 向服务层请求数据
	messageResponse, err := service.MessageAction(tokenUid, toUserId, content, actionType)
	if err != nil {
		fmt.Printf("message action error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", messageResponse)

}

func MessageChatList(ctx *gin.Context) {

	var err error
	tokenUids, _ := ctx.Get("UserId")
	tokenUid := tokenUids.(int64)
	to_user_id := ctx.Query("to_user_id")
	toUserId, err := strconv.ParseInt(to_user_id, 10, 64)
	if err != nil {
		fmt.Printf("handler func-MessageChatList toUserId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 得到上次请求的最新时间
	preMsgTimes := ctx.Query("pre_msg_time")
	pre_msg_time, err := strconv.ParseInt(preMsgTimes, 10, 64)
	if err != nil {
		fmt.Println("parse preMsgTime error!")
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 向服务层请求数据
	messageListResponse, err := service.MessageList(tokenUid, toUserId, pre_msg_time)
	if err != nil {
		fmt.Printf("messageListResponse error : %s", err)
		// log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", messageListResponse)
}
