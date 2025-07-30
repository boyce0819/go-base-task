package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-task1/bolg/errorss"
	"go-task1/bolg/results"
	"gorm.io/gorm"
	"net/http"
)

func GlobalErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 检查是否有错误
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				// 处理不同类型的错误
				switch {
				case errors.Is(err.Err, errorss.ErrInvalidCredentials):
					results.FailStop(c, errorss.ErrInvalidCredentials.Code, errorss.ErrInvalidCredentials.Error())
					return
				case errors.Is(err.Err, errorss.ErrInvalidParams):
					results.FailStop(c, errorss.ErrInvalidParams.Code, errorss.ErrInvalidParams.Error())
					return
				case errors.Is(err.Err, errorss.ErrUnauthorized):
					results.FailStop(c, errorss.ErrUnauthorized.Code, errorss.ErrUnauthorized.Error())
					return
				case errors.Is(err.Err, gorm.ErrRecordNotFound):
					results.FailStop(c, http.StatusInternalServerError, "数据不存在")
					return
				default:
					// 默认错误处理
					results.FailStop(c, errorss.ErrSystem.Code, errorss.ErrSystem.Message)
					return
				}
			}
		}
	}
}
