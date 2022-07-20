package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	i := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(i); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, models.CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, models.CodeInvalidParam, errData)
		return
	}
	service.InvitationVote()
	ResponseSuccess(c, nil)
}
