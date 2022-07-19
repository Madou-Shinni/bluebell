package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/models"
	"web_app/service"
)

// 社区

// CommunityHandler 社区
func CommunityHandler(c *gin.Context) {
	// 1.查询社区列表
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, models.CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 根据id获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, models.CodeInvalidParam)
		return
	}
	// 2.查询社区
	data, err := service.GetCommunityDetailById(id)
	if err != nil {
		ResponseError(c, models.CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
