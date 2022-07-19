package controller

import "github.com/gin-gonic/gin"

// AddInvitationHandler 添加帖子
func AddInvitationHandler(c *gin.Context) {
	// 1.获取参数
	c.ShouldBindJSON()
	// 2.添加帖子

	// 3.响应
}
