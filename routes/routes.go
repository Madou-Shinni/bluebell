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
		v1.GET("/invitation/add", controller.AddInvitationHandler)
		v1.GET("/invitation/:id", controller.GetInvitationDetailHandler)
	}
	return r
}
