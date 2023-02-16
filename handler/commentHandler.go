package handler

import (
	"fmt"
	"strconv"
	"tiktok/response"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

// 发布评论
func CommentAction(ctx *gin.Context) {
	var err error

	tokenUids, _ := ctx.Get("UserId")

	tokenUid := tokenUids.(int64)

	video_id := ctx.Query("video_id")
	comment_text := ctx.Query("comment_text")
	actionType := ctx.Query("action_type")
	comment_id := ctx.Query("comment_id")
	commentId := int64(0)
	if actionType == "2" {
		commentId, err = strconv.ParseInt(comment_id, 10, 64)
		if err != nil {
			fmt.Printf("commentId error : %s", err)
			response.Fail(ctx, err.Error(), nil)
			return
		}
	}
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		fmt.Printf("videoId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	commentResponse, err := service.CommentAction(commentId, videoId, tokenUid, comment_text, actionType)
	if err != nil {
		fmt.Printf("comment error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", commentResponse)
}

// 获取评论列表
func GetCommentList(ctx *gin.Context) {
	var err error
	video_id := ctx.Query("video_id")
	/* token := ctx.Query("token")
		_, err = util.VerifyToken(token)
		if err != nil {
			log.Errorf("token error : %s", err)
	   		response.Fail(ctx, err.Error(), nil)
			return
		} */
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		fmt.Printf("videoId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	listResponse, err := service.CommentList(videoId)
	if err != nil {
		fmt.Printf("list error : %s", err)
		// log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", listResponse)
}
