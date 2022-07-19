package service

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/tools/snowflake"
)

// AddInvitation 添加帖子
func AddInvitation(i *models.Invitation) (err error) {
	// 1.生成id
	i.Id = snowflake.GenID()
	// 2.保存到mysql
	return mysql.InsertInvitation(i)
}
