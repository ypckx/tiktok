package base_handler

import (
	"net/http"
	"tinyTiktok/services/base_service"

	"github.com/gin-gonic/gin"
)

/***用户注册请求
[method]: POST
[路由]: /douyin/user/register
[请求参数]:
参数名	     位置	 类型	必填	说明
username	query	string  是   注册用户名，最长32个字符
password	query	string  是   密码，最长32个字符

响应：
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
***/

func RegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	registerResp, _ := base_service.UserReigsterService(username, password)

	c.JSON(http.StatusOK, registerResp)
}
