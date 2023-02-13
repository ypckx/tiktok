package main

import (
	"fmt"
	"tiktok/config"
	"tiktok/model/db"
	"tiktok/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	// 静态文件访问路径
	r.Static("/static", "./public")
	r = router.InitRouter(r)

	r.Run(config.ServerIp)

	fmt.Println("after gin.Default()... after run....")
}
