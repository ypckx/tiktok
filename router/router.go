package router

import (
	"tinyTiktok/handlers/base_handler"
	"tinyTiktok/handlers/interact_handler"
	"tinyTiktok/handlers/social_handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", base_handler.FeedHandler)
	apiRouter.GET("/user/", base_handler.UserInfoHandler)
	apiRouter.POST("/user/register/", base_handler.RegisterHandler)
	apiRouter.POST("/user/login/", base_handler.LoginHandler)
	apiRouter.POST("/publish/action/", base_handler.PublishHandler)
	apiRouter.GET("/publish/list/", base_handler.PublishListHandler)

	// extra apis - I
	apiRouter.POST("/favorite/action/", interact_handler.FavoriteActionHandler)
	apiRouter.GET("/favorite/list/", interact_handler.FavoriteListHandler)
	apiRouter.POST("/comment/action/", interact_handler.CommentActionHandler)
	apiRouter.GET("/comment/list/", interact_handler.CommentListHandler)

	// extra apis - II
	apiRouter.POST("/relation/action/", social_handler.RelationActionHandler)
	apiRouter.GET("/relation/follow/list/", social_handler.FollowListHandler)
	apiRouter.GET("/relation/follower/list/", social_handler.FollowerListHandler)
	apiRouter.GET("/relation/friend/list/", social_handler.FriendListHandler)
	apiRouter.GET("/message/chat/", social_handler.MessageChatHandler)
	apiRouter.POST("/message/action/", social_handler.MessageActionHandler)
}
