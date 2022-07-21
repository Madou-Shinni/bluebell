package redis

import "web_app/models"

// GetInvitationIdsInOrder 去Redis查询id列表
func GetInvitationIdsInOrder(p *models.ParamInvitationList) ([]string, error) {
	// 1.根据用户请求中携带的参数来决定根据什么排序
	key := getRedisKey(KeyInvitationTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyInvitationTimeZSet)

	}
	// 2.查询的索引起始点终点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.ZREVRANG 查询，大——>小排序
	return rdb.ZRevRange(key, start, end).Result()
}
