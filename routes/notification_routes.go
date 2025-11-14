package routes

import (
	"dinsos_kuburaya/controllers"
	"dinsos_kuburaya/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine) {
	api := r.Group("/api")

	auth := api.Group("/notifications")
	auth.Use(middleware.RequireAuth()) // <-- Lindungi
	{
		auth.GET("", controllers.GetNotifications)
		auth.POST("/:id/read", controllers.MarkNotificationAsRead)
	}
}
