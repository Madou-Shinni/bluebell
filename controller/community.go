package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
