package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/models"
	"web_app/service"
)

// SignUpHandler 注册处理
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok { // 如果不是
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误!",
			})
			return
		}
		// 如果是就把错误翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误，并去除提示信息中的结构体名称
		})
		return
	}
	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败！"})
		return
	}
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok { // 如果不是
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误!",
			})
			return
		}
		// 如果是就把错误翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误，并去除提示信息中的结构体名称
		})
		return
	}
	// 2.业务逻辑处理
	if err := service.Login(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名或密码错误！"})
		return
	}
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
