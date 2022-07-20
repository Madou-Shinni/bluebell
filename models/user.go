package models

type User struct {
	UserId   int64  `json:"userId,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Token    string `json:"token"`
}
