package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/models"
	"web_app/service"
)

// VoteData 投票结构体
type VoteData struct {
	// userId 从当前用户中获取
	// 帖子id
	InvitationId int64 `json:"invitationId,string" binding:"required"`
	// 赞成票(1) 还是 反对票(-1) 取消投票(0)
	Direction int8 `json:"direction,string" binding:"required,oneof=1 -1 0"`
}

// InvitationVoteHandler 投票处理
func InvitationVoteHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, models.CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, models.CodeInvalidParam, errData)
		return
	}
	// 获取当前请求的用户id
	userId, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, models.CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := service.InvitationVote(userId, p); err != nil {
		zap.L().Error("service.InvitationVote() failed", zap.Error(err))
		ResponseError(c, models.CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
