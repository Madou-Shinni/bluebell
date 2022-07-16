package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"web_app/tools/jwt"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokenStr := c.Request.Header.Get("Authorization")
		// 用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		// token格式错误
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token格式错误"})
			c.Abort() //阻止执行
			return
		}
		// 解析token
		tokenStruck, err := jwt.ParseToken(tokenSlice[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token不正确"})
			c.Abort() //阻止执行
			return
		}
		// token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token过期"})
			c.Abort() //阻止执行
			return
		}
		c.Set("userId", tokenStruck.UserId)
		c.Set("username", tokenStruck.Username)

		c.Next()
	}
}
