package models

// Community 社区
type Community struct {
	Id   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
