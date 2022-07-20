package models

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
	Direction int8 `json:"direction,string" binding:"required,oneof=1 -1 0"`
}
