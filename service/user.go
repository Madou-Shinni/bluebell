package service

import (
	"web_app/dao/mysql"
	snowflake "web_app/tools"
)

// 用户注册
func SignUp() {
	// 1.判断用户名是否存在
	mysql.QueryUserByUsername()
	// 2.生成ID
	snowflake.GenID()
	// 3.密码加密
	// 4.添加用户
	mysql.InsertUser()
}