package service

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetCommunityList 查询社区列表
func GetCommunityList() ([]*models.Community, error) {
	return mysql.SelectCommunityList()
}
