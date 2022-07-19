package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

// GetInvitationDetailHandler 查询帖子详情
func GetInvitationDetailHandler(c *gin.Context) {
	// 1.获取参数
	iIdStr := c.Param("id")
	iId, err := strconv.ParseInt(iIdStr, 10, 64)
	if err != nil {
		zap.L().Error("Get Invitation Detail with invalid param failed", zap.Error(err))
		return
	}
	// 2.根据id查询帖子
	data, err := service.GetInvitationDetailById(iId)
	if err != nil {
		zap.L().Error("service.GetInvitationDetailById(iId) failed", zap.Error(err))
		ResponseError(c, models.CodeServerBusy)
		return
	}
	// 3.响应
	ResponseSuccess(c, data)
}

// GetInvitationListHandler 获取帖子列表
func GetInvitationListHandler(c *gin.Context) {
	// 1.获取分页参数
	pageNum, pageSize := GetPageInfo(c)
	// 2.查询帖子列表
	data, err := service.GetInvitationList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("service.GetInvitationList() failed", zap.Error(err))
		ResponseError(c, models.CodeServerBusy)
		return
	}
	// 3.响应
	ResponseSuccess(c, data)
}
