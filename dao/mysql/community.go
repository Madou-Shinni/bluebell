package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

// SelectCommunityList 查询社区列表
func SelectCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows { // 没有记录
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// SelectCommunityDetailById 社区id查询社区详情
func SelectCommunityDetailById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time,update_time from community where community_id = ?`
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows { // 没有记录
			zap.L().Warn("there is no communityDetail in db")
			err = ErrorInvalidId
		}
	}
	return community, err
}
