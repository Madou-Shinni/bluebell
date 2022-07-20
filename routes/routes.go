package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/logger"
	"web_app/middleware"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
	}
	v1.Use(middleware.JwtMiddleware()) // 应用认证中间件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

	}

	{
		v1.POST("/invitation/add", controller.AddInvitationHandler)
		v1.GET("/invitation/:id", controller.GetInvitationDetailHandler)
		v1.GET("/invitation/list", controller.GetInvitationListHandler)
		v1.POST("/invitation/vote", controller.InvitationVoteHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/invitation/listBy", controller.GetInvitationListByHandler)
	}
	return r
}
