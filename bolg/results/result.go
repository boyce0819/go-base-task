package results

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/errorss"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Result{
		Code:    200,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
func Error(c *gin.Context, appErr *errorss.AppError) {
	c.JSON(http.StatusOK, Result{
		Code:    appErr.Code,
		Message: appErr.Message,
		Data:    nil,
	})
}

func FailStop(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, Result{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
