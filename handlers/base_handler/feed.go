package base_handler

import (
	"net/http"
	"tinyTiktok/services/base_service"

	"github.com/gin-gonic/gin"
)

// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

func FeedHandler(c *gin.Context) {
	latest_time := c.Query("latest_time")
	userToken := c.Query("token")
	feedResp, _ := base_service.FeedService(latest_time, userToken)
	c.JSON(http.StatusOK, feedResp)
}
