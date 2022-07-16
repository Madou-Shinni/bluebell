package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"web_app/controller"
	"web_app/models"
	"web_app/tools/jwt"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokenStr := c.Request.Header.Get("Authorization")
		// 用户不存在
		if tokenStr == "" {
			controller.ResponseError(c, models.CodeNeedLogin)
			c.Abort() //阻止执行
			return
		}
		// token格式错误
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			controller.ResponseError(c, models.CodeInvalidToken)
			c.Abort() //阻止执行
			return
		}
		// 解析token
		tokenStruck, err := jwt.ParseToken(tokenSlice[1])
		if err != nil {
			controller.ResponseError(c, models.CodeInvalidToken)
			c.Abort() //阻止执行
			return
		}
		// token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			controller.ResponseError(c, models.CodeExpireToken)
			c.Abort() //阻止执行
			return
		}
		// 将请求的userId信息保存到上下文，后续可以通过c.Get("userId")获取
		c.Set(controller.ContextUserIdKey, tokenStruck.UserId)
		c.Next()
	}
}
