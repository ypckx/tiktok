package handler

import (
	"fmt"
	"strconv"
	"tiktok/response"
	"tiktok/service"
	"tiktok/utils"

	"github.com/gin-gonic/gin"
)

func FeedHandler(c *gin.Context) {
	fmt.Println("FeedHandler begin==============")
	// return
	var userId int64

	currentTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	if err != nil || currentTime == 0 {
		currentTime = utils.GetCurrentTime()
	}

	userIds, _ := c.Get("UserId")
	userId = userIds.(int64)

	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	feedList, err := service.GetFeedList(currentTime, userId)
	// fmt.Println("FeedHandler end========== feedList", feedList)

	if err != nil {
		response.Fail(c, err.Error(), nil)
	}
	response.Success(c, "success", feedList)
}
