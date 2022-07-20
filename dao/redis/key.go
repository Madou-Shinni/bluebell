package redis

// redis key

const (
	KeyPrefix                    = "bluebell:"
	KeyInvitationTimeZSet        = "post:time"   // 帖子及发帖时间
	KeyInvitationScoreZSet       = "post:score"  // 帖子及投票的分数
	KeyInvitationVotedZSetPrefix = "post:voted:" // 记录用户及投票的类型（不完整的key）
)
