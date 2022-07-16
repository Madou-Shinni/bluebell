package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ContextUserIdKey = "userId" // userId
var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前用户Id
func GetCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	if userId, ok = uid.(int64); !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}
