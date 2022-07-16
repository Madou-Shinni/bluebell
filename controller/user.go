package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
			ResponseError(c, models.CodeInvalidParam)
			return
		}
		// 如果是就把错误翻译
		ResponseErrorWithMsg(c, models.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		if errors.Is(err, service.ErrorUserExist) { // 比较抛出的错误信息
			ResponseError(c, models.CodeUserExist)
			return
		}
		ResponseErrorWithMsg(c, models.CodeServerBusy, models.CodeServerBusy.Msg())
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
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
			ResponseError(c, models.CodeInvalidParam)
			return
		}
		// 如果是就把错误翻译
		ResponseErrorWithMsg(c, models.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务逻辑处理
	token, err := service.Login(p)
	if errors.Is(err, service.ErrorPassword) { // 比较抛出的错误信息
		ResponseError(c, models.CodePassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, token)
}
