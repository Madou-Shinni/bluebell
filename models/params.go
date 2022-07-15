package models

// 请求参数结构体
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}
