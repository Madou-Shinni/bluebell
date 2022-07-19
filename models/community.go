package models

import "time"

// Community 社区
type Community struct {
	Id   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// CommunityDetail 社区详情
type CommunityDetail struct {
	Id           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
	UpdateTime   time.Time `json:"updateTime" db:"update_time"`
}
