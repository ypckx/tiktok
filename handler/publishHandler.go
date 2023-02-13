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
	// publishResponse := &message.DouyinPublishActionResponse{}
	userId, _ := ctx.Get("UserId")
	//token := ctx.PostForm("token")
	//userId, err := common.VerifyToken(token)
	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%s_%s", utils.RandomString(), filename)
	// videoPath := config.GetConfig().Path.Videofile
	videoPath := config.VideoPath
	saveFile := filepath.Join(videoPath, finalName)

	// log.Info("saveFile:", saveFile)
	// fmt.Println("[PublishAction] ==== saveFile:", saveFile)

	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		fmt.Println("ctx.SaveUploadedFile err:", err.Error())
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// fmt.Println("[publishHandler.go -> func-PublishAction] userId:", userId, " title:", title, " savePath:", saveFile)
	// return
	publish, err := service.PublishVideo(userId.(int64), saveFile, title)
	//publish, err := service.PublishVideo(userId, saveFile)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	// log.Infof("publish:%v", publish)
	fmt.Printf("publish:%v", publish)
	response.Success(ctx, "success", publish)

}

func GetPublishList(ctx *gin.Context) {

	tokenUserId, _ := ctx.Get("UserId")
	id := ctx.Query("user_id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}

	// fmt.Println("ctx.Query(user_id):", id)

	// fmt.Println("[publishHandler.go func-GetPublishList] tokenUserId:", tokenUserId, "  userId:", userId)
	list, err := service.PublishList(tokenUserId.(int64), userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", list)
}
