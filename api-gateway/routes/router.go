package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "pong")
		})
		userGroup := v1.Group("/user")
		userGroup.POST("/register", handler.UserRegisterHandler)
		userGroup.POST("/login", handler.UserLoginHandler)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 任务模块
			authed.GET("task", handler.GetTaskList)
			authed.POST("task", handler.CreateTask)
			authed.PUT("task", handler.UpdateTask)
			authed.DELETE("task", handler.DeleteTask)
		}
	}
	return ginRouter
}
