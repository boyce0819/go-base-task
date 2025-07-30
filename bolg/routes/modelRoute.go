package routes

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/controller"
	"go-task1/bolg/middleware"
)

func ApiRouter() *gin.Engine {
	r := gin.Default()
	// 不需要认证的路由
	public := r.Group("/api")
	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
		//public.GET("/health", controller.HealthCheck)
	}

	// 需要JWT认证的路由
	protected := r.Group("/api")
	protected.Use(
		middleware.AuthMiddleware(),
		middleware.GlobalErrorHandlerMiddleware()) // 应用JWT中间件  全局异常中间件 todo
	{
		protected.POST("/post/create", controller.CreatePost)
		protected.POST("/post/detail", controller.GetPostDetail)
		protected.POST("/post/list", controller.ListPost)
		protected.POST("/post/delete", controller.DeletePost)
		protected.POST("/post/update", controller.UpdatePost)
		protected.POST("/comment/create/", controller.CommentCreate)
		protected.POST("/commentPost/list", controller.CommentPostList)
	}
	return r
}
