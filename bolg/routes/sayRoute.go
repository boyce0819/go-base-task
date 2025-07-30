package routes

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/controller"
)

func SayRouter() *gin.Engine {
	router := gin.Default()
	// 注册路由
	router.GET("/say", controller.SeyHello)
	router.GET("/goodbye", controller.SayGoodbye)

	return router
}
