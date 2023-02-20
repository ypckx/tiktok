package config

// 服务器运行地址
var ServerIp string = "100.92.33.215:8080"

// 聊天服务器地址
var MessageServerIp string = "100.92.33.215:8080"

/********数据库配置信息**************

数据库使用Gorm连接：https://gorm.io/zh_CN/docs/index.html
数据库类型：mysql
登录用户：	gorm
密码：		abc123
数据库地址：localhost:3306
数据库名称：tiktok
数据库字符集：charset=utf8mb4
parseTime=True 确保正确处理time.Time类型
**********************************/

var Dns string = "gorm:abc123@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

// 视频url地址
var VideoUrl string = "http://" + ServerIp + "/static/video/"
var CoverUrl string = "http://" + ServerIp + "/static/pic/"

var basePath string = "E:/CodeSource/Golang_code/tiktok"
var VideoPath string = basePath + "/public/video/"
var PicPath string = basePath + "/public/pic/"
