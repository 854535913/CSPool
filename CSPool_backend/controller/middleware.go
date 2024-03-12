package controller

import (
	"github.com/gin-gonic/gin"
	"main/service"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ResponseError(c, CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ResponseError(c, CodeTokenFormatError)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		info, err := service.DecryptTokenService(parts[1])
		if err != nil {
			ResponseError(c, CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文c上
		c.Set("UserID", info.UserID)
		c.Set("Username", info.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
