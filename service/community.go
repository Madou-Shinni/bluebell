package service

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetCommunityList 查询社区列表
func GetCommunityList() ([]*models.Community, error) {
	return mysql.SelectCommunityList()
}

// GetCommunityDetailById 根据id查询社区
func GetCommunityDetailById(id int64) (*models.CommunityDetail, error) {
	return mysql.SelectCommunityDetailById(id)
}
