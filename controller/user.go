package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app/models"
	"web_app/service"
)

// 注册处理
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误!",
		})
		return
	}
	// 手动参数校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword == p.Password {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误!",
		})
		return
	}
	// 2.业务处理
	service.SignUp(p)
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
