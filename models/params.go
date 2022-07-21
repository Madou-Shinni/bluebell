package models

const (
	OrderTime  = "time"  // 按时间排序
	OrderScore = "score" // 按分数排序
)

// 请求参数结构体

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数
type ParamVoteData struct {
	// userId 从当前用户中获取
	// 帖子id
	InvitationId int64 `json:"invitationId,string" binding:"required"`
	// 赞成票(1) 还是 反对票(-1) 取消投票(0)
	Direction int8 `json:"direction,string" binding:"oneof=1 -1 0"`
}

// ParamInvitationList 获取帖子列表参数
type ParamInvitationList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size"form:"size"`
	Order string `json:"order" form:"order"`
}
