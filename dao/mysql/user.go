package mysql

import "web_app/models"

// QueryUserByUsername 根据用户名查询用户
func QueryUserByUsername(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	return
}

// InsertUser 向数据库中添加一条用户记录
func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	if err != nil {
		return err
	}
	return
}
