package middleware

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/errorss"
	"go-task1/bolg/utils"
	"strings"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errorss.ThrowErr(c, errorss.ErrInvalidCredentials, "请求头中没有找到token")
			return
		}

		// 解析Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errorss.ThrowErr(c, errorss.ErrInvalidCredentials, "token格式有问题")
			return
		}

		// 验证Token
		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			errorss.ThrowErr(c, errorss.ErrInvalidCredentials, "token无效： "+err.Error())
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.UserName)
		c.Next()
	}
}
