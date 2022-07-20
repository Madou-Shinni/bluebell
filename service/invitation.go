package service

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/tools/snowflake"
)

// AddInvitation 添加帖子
func AddInvitation(i *models.Invitation) (err error) {
	// 1.生成id
	i.InvitationId = snowflake.GenID()
	// 2.保存到mysql
	err = mysql.InsertInvitation(i)
	if err != nil {
		return err
	}
	// 3.保存到redis
	err = redis.AddInvitation(i.InvitationId)
	if err != nil {
		return err
	}
	return
}

// GetInvitationDetailById 根据帖子id获取帖子详情
func GetInvitationDetailById(iId int64) (data *models.ApiInvitationDetail, err error) {
	// 查询数据并组合我们接口想用的数据集
	i, err := mysql.SelectInvitationById(iId)
	if err != nil {
		zap.L().Error("mysql.SelectInvitationById(iId) failed", zap.Error(err))
		return
	}
	// 根据作者Id查询作者信息
	user, err := mysql.GetUserById(i.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(i.AuthorId) failed",
			zap.Int64("authorId", i.AuthorId),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详情
	c, err := mysql.SelectCommunityDetailById(i.CommunityId)
	if err != nil {
		zap.L().Error("mysql.SelectCommunityDetailById(i.CommunityId) failed",
			zap.Int64("communityId", i.CommunityId),
			zap.Error(err))
		return
	}
	data = &models.ApiInvitationDetail{
		AuthorName:      user.Username,
		Invitation:      i,
		CommunityDetail: c,
	}
	return
}

// GetInvitationList 获取帖子列表
func GetInvitationList(pageNum, pageSize int64) (data []*models.ApiInvitationDetail, err error) {
	list, err := mysql.SelectInvitationList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("mysql.SelectInvitationList() failed", zap.Error(err))
		return
	}
	data = make([]*models.ApiInvitationDetail, 0, len(list))
	for _, i := range list {
		// 根据作者Id查询作者信息
		user, err := mysql.GetUserById(i.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(i.AuthorId) failed",
				zap.Int64("authorId", i.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详情
		c, err := mysql.SelectCommunityDetailById(i.CommunityId)
		if err != nil {
			zap.L().Error("mysql.SelectCommunityDetailById(i.CommunityId) failed",
				zap.Int64("communityId", i.CommunityId),
				zap.Error(err))
			continue
		}
		apiInvitationDetail := &models.ApiInvitationDetail{
			AuthorName:      user.Username,
			Invitation:      i,
			CommunityDetail: c,
		}
		data = append(data, apiInvitationDetail)
	}

	return
}
