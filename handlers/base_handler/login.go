package base_handler

import (
	"net/http"
	"tinyTiktok/services/base_service"

	"github.com/gin-gonic/gin"
)

/***用户注册请求
[method]: POST
[路由]: /douyin/user/register
[说明]: 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token

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

// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
func LoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	loginResp, _ := base_service.UserLoginService(username, password)
	c.JSON(http.StatusOK, loginResp)
}
