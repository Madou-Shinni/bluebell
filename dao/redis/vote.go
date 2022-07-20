package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"time"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
   direction=1时，有两种情况：
   	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2  +432*2
   direction=0时，有两种情况：
   	1. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  -432
   direction=-1时，有两种情况：
   	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录  差值的绝对值：2  -432*2

   投票的限制：
   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
   	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
   	2. 到期之后删除那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600 // 一周
	scorePreVote     = 432           // 每一票占的分数
)

var (
	ErrorVoteTimeExpire = errors.New("超出投票时间")
)

// AddInvitation 添加帖子
func AddInvitation(invitationId int64) error {
	pipeline := rdb.TxPipeline() // 添加事务命令
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyInvitationTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: invitationId,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyInvitationScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: invitationId,
	})
	_, err := pipeline.Exec()
	if err != nil {
		zap.L().Error("redis AddInvitation failed", zap.Error(err))
	}
	return err
}

// InvitationVote 帖子投票
func InvitationVote(userId, invitationId string, value float64) error {
	// 1.判断投票限制
	// 从redis中去投票时间
	invitationTime := rdb.ZScore(getRedisKey(KeyInvitationTimeZSet), invitationId).Val()
	if float64(time.Now().Unix())-invitationTime > oneWeekInSeconds { // 当前时间 - 投票时间 > 一周
		return ErrorVoteTimeExpire
	}

	// 2和3需要加入事务

	// 2.更新帖子分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyInvitationVotedZSetPrefix+invitationId), userId).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值

	pipeline := rdb.TxPipeline() // 事务
	pipeline.ZIncrBy(getRedisKey(KeyInvitationScoreZSet), dir*diff*scorePreVote, invitationId)
	// 3.记录用户为该帖子投票的数据
	if value == 0 { // 取消投票（移除key）
		pipeline.ZRem(getRedisKey(KeyInvitationVotedZSetPrefix), userId)
	} else {
		pipeline.ZAdd(getRedisKey(KeyInvitationVotedZSetPrefix+invitationId), redis.Z{
			Score:  value, // 当前赞成票还是反对票
			Member: userId,
		})
	}
	_, err := pipeline.Exec()
	return err
}
