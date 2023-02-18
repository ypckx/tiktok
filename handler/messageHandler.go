package handler

import (
	"fmt"
	"strconv"
	"tiktok/response"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

// messageRouter.POST("/action", common.AuthMiddleware(), handler.MessageAction)
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
	//==================
	fmt.Println("[handler MessageAction] =======> tokenUid:", toUserId, " to_user_id:", toUserId, " content:", content)

	messageResponse, err := service.MessageAction(tokenUid, toUserId, content, actionType)

	if err != nil {
		fmt.Printf("message action error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", messageResponse)

}

// messageRouter.GET("/chat", common.AuthWithOutMiddleware(), handler.MessageChatList)
func MessageChatList(ctx *gin.Context) {
	// return
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

	// add key-value to map struct

	//==================
	// fmt.Println("[handler MessageChatList]========> tokenUid:", tokenUid, " to_user_id:", to_user_id)

	messageListResponse, err := service.MessageList(tokenUid, toUserId)
	if err != nil {
		fmt.Printf("messageListResponse error : %s", err)
		// log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", messageListResponse)
}
