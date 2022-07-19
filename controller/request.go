package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
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

// GetPageInfo 获取分页参数
func GetPageInfo(c *gin.Context) (int64, int64) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageNum")
	var (
		pageNum  int64
		pageSize int64
		err      error
	)
	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	return pageNum, pageSize
}
