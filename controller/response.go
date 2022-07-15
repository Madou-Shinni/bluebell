package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/models"
)

// 封装响应结构体、响应函数

// ResponseData 同一响应数据格式
type ResponseData struct {
	Code models.ResCode `json:"code"`
	Msg  interface{}    `json:"msg"`
	Data interface{}    `json:"data"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code models.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 响应错误及自定义提示信息
func ResponseErrorWithMsg(c *gin.Context, code models.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: models.CodeSuccess,
		Msg:  models.CodeSuccess.Msg(),
		Data: data,
	})
}
