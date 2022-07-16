package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/tools/jwt"
	"web_app/tools/snowflake"
)

var (
	ErrorUserExist    = errors.New("用户名已存在")
	ErrorUserNotExist = errors.New("用户名不存在")
	ErrorPassword     = errors.New("用户名或密码错误")
)

// 密码盐
var salt = "www.yumclor.com"

// SignUp 用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户名是否存在
	if err = mysql.QueryUserByUsername(p.Username); err != nil {
		zap.L().Error("SignUp Error", zap.Error(err))
		return ErrorUserNotExist
	}
	// 2.生成ID
	userId := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: md5Password(p.Password), // 3.密码加密
	}
	// 4.添加用户
	if err = mysql.InsertUser(user); err != nil {
		zap.L().Error("SignUp Error", zap.Error(err))
		return err
	}
	return
}

// Login 登录
func Login(p *models.ParamLogin) (token string, err error) {
	oldPassword := md5Password(p.Password)
	// 构造user实例
	user := &models.User{
		Username: p.Username, Password: md5Password(p.Password),
	}
	// 通过用户名密码查询用户（参数传递的是指针）
	if err = mysql.SelectUserByUsername(user); err != nil {
		zap.L().Error("Login Error", zap.Error(err))
		return "", ErrorUserNotExist
	}
	// 判断密码是否正确
	if oldPassword != user.Password {
		return "", ErrorPassword
	}
	// 生成jwt的token
	return jwt.GenToken(user.UserId, user.Username)
}

// md5Password 加密
func md5Password(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(salt))                                 // 加密码盐
	return hex.EncodeToString(h.Sum([]byte(oldPassword))) // 把字节数组转换成16进制的字符串
}
