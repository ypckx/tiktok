package handler

import (
	"fmt"
	"path/filepath"
	"strconv"
	"tiktok/config"
	"tiktok/response"
	"tiktok/service"
	"tiktok/utils"

	"github.com/gin-gonic/gin"
)

// 视频发布
func PublishAction(ctx *gin.Context) {

	userId, _ := ctx.Get("UserId")
	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 产生视频文件名
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s", utils.RandomString(), filename)
	videoPath := config.VideoPath
	saveFile := filepath.Join(videoPath, finalName)

	// 文件保存
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		fmt.Println("ctx.SaveUploadedFile err:", err.Error())
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 向服务层请求数据
	publish, err := service.PublishVideo(userId.(int64), saveFile, title)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", publish)

}

func GetPublishList(ctx *gin.Context) {

	tokenUserId, _ := ctx.Get("UserId")
	id := ctx.Query("user_id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}

	// 向服务层请求数据
	list, err := service.PublishList(tokenUserId.(int64), userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", list)
}
