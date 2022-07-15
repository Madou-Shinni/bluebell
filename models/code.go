package models

// 同一返回数据代码

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodePassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist:    "用户名已存在",
	CodeUserNotExist: "用户名不存在",
	CodePassword:     "用户名或密码错误",
	CodeServerBusy:   "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
