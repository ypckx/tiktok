package router

import (
	"tiktok/common"
	"tiktok/handler"

	"github.com/gin-gonic/gin"
)

func setUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	userRouter.POST("/login/", handler.UserLogin)
	userRouter.GET("/", common.AuthMiddleware(), handler.GetUserInfo)
	userRouter.POST("/register/", handler.UserRegister)
}

func setPublishRouter(r *gin.RouterGroup) {
	publishRouter := r.Group("publish")
	publishRouter.POST("/action/", common.AuthMiddleware(), handler.PublishAction)
	publishRouter.GET("/list/", common.AuthWithOutMiddleware(), handler.GetPublishList)
}

func setCommentRouter(r *gin.RouterGroup) {
	commentRouter := r.Group("comment")
	commentRouter.POST("/action/", common.AuthMiddleware(), handler.CommentAction)
	commentRouter.GET("/list/", common.AuthWithOutMiddleware(), handler.GetCommentList)
}

func setFavoriteRouter(r *gin.RouterGroup) {
	favoriteRouter := r.Group("favorite")
	favoriteRouter.POST("/action/", common.AuthMiddleware(), handler.FavoriteAction)
	favoriteRouter.GET("/list/", common.AuthWithOutMiddleware(), handler.GetFavoriteList)
}

func setRelationRouter(r *gin.RouterGroup) {
	relationRouter := r.Group("relation")
	relationRouter.POST("/action/", common.AuthMiddleware(), handler.RelationAction)
	relationRouter.GET("/follow/list", common.AuthWithOutMiddleware(), handler.GetFollowList)
	relationRouter.GET("/follower/list", common.AuthWithOutMiddleware(), handler.GetFollowerList)
}

func InitRouter(r *gin.Engine) *gin.Engine {
	apiRouter := r.Group("/douyin")

	setUserRouter(apiRouter)
	setPublishRouter(apiRouter)
	setCommentRouter(apiRouter)
	setFavoriteRouter(apiRouter)
	setRelationRouter(apiRouter)

	apiRouter.GET("/feed/", common.AuthWithOutMiddleware(), handler.FeedHandler)

	return r
}
