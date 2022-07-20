package redis

// redis key

const (
	KeyPrefix                    = "bluebell:"
	KeyInvitationTimeZSet        = "invitation:time"   // 帖子及发帖时间
	KeyInvitationScoreZSet       = "invitation:score"  // 帖子及投票的分数
	KeyInvitationVotedZSetPrefix = "invitation:voted:" // 记录用户及投票的类型（不完整的key）
)

// getRedisKey 得到redisKey
func getRedisKey(key string) string {
	return KeyPrefix + key
}
