package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/models"
	"web_app/service"
)

// AddInvitationHandler 添加帖子
func AddInvitationHandler(c *gin.Context) {
	// 1.获取参数
	i := new(models.Invitation)
	if err := c.ShouldBindJSON(i); err != nil {
		zap.L().Debug("c.ShouldBindJSON(i) err", zap.Any("err", err))
		zap.L().Error("Add Invitation with invalid param failed", zap.Error(err))
		ResponseError(c, models.CodeInvalidParam)
		return
	}
	// 从c中获取当前用户id
	userId, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, models.CodeNeedLogin)
		return
	}
	i.AuthorId = userId
	// 2.添加帖子
	if err := service.AddInvitation(i); err != nil {
		zap.L().Error("service.AddInvitation failed", zap.Error(err))
		ResponseError(c, models.CodeServerBusy)
		return
	}
	// 3.响应
	ResponseSuccess(c, models.CodeSuccess)
}
