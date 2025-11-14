package routes

import (
	"dinsos_kuburaya/controllers"
	"dinsos_kuburaya/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Login endpoint (tidak perlu auth)
	api.POST("/login", controllers.Login)

	// Rute yang memerlukan autentikasi
	auth := api.Group("/auth")
	auth.Use(middleware.RequireAuth()) // Terapkan middleware
	{
		auth.GET("/me", controllers.GetMe)
		auth.POST("/logout", controllers.Logout)
	}
}
