package models

import "time"

// Invitation 帖子
type Invitation struct { // go 内存对齐（相同类型的字段放在一起可以减少内存占用）
	Id           int64     `json:"id" db:"id"`
	InvitationId int64     `json:"invitationId,string" db:"invitation_id"`
	AuthorId     int64     `json:"authorId,string" db:"author_id"`
	CommunityId  int64     `json:"communityId" db:"community_id" binding:"required"`
	Status       int32     `json:"status" db:"status"`
	Title        string    `json:"title" db:"title" binding:"required"`
	Content      string    `json:"content" db:"content"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
	UpdateTime   time.Time `json:"updateTime" db:"update_time"`
}

// ApiInvitationDetail 帖子详情接口结构体
type ApiInvitationDetail struct {
	AuthorName       string                   `json:"authorName"`
	*Invitation      `json:"invitation"`      // 嵌入帖子结构体
	*CommunityDetail `json:"communityDetail"` // 嵌入社区结构体
}
