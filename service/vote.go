package service

import (
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

// 投票:1.用户投票的数据2.

// InvitationVote 帖子投票
func InvitationVote(userId int64, p *models.ParamVoteData) error {
	zap.L().Debug("InvitationVote",
		zap.Int64("userId", userId),
		zap.Int64("InvitationId", p.InvitationId), zap.Int8("Direction", p.Direction))
	return redis.InvitationVote(
		strconv.FormatInt(userId, 10),
		strconv.FormatInt(p.InvitationId, 10),
		float64(p.Direction))
}
