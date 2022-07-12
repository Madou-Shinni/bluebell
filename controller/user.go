package controller

import (
	"github.com/gin-gonic/gin"
	"web_app/service"
)

// 注册处理
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	// 2.业务处理
	service.SignUp()
	// 3.返回响应
}
