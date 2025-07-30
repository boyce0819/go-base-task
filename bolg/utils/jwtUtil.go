package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte("your-secret-key") // 生产环境应从环境变量获取

type Claims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(userID int, userName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	claims := &Claims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    "jwt-auth-go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义不需要验证的路由
		skipPaths := map[string]bool{
			"/api/register": true,
			"/api/login":    true,
			"/api/health":   true,
		}
		// 如果当前路径在跳过列表中，直接进入下一个处理程序
		if skipPaths[c.FullPath()] {
			c.Next()
			return
		}
		// 以下是常规JWT验证逻辑
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
	}
}
