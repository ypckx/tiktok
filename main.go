package main

import (
	"tinyTiktok/models"
	"tinyTiktok/myconfig"
	"tinyTiktok/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.InitDB()

	router.InitRouter(r)

	r.Run(myconfig.ServerIp)
}
