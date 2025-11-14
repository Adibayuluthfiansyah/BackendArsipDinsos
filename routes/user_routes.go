package routes

import (
	"dinsos_kuburaya/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)       // → /api/users
		api.GET("/users", controllers.GetUsers)          // → /api/users
		api.GET("/users/:id", controllers.GetUserByID)   // → /api/users/:id
		api.PUT("/users/:id", controllers.UpdateUser)    // → /api/users/:id
		api.DELETE("/users/:id", controllers.DeleteUser) // → /api/users/:id
	}
}
