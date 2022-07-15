package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/logger"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controller.SignUpHandler)
	return r
}
