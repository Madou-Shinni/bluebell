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
