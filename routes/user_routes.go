package routes

import (
	"dinsos_kuburaya/controllers"
	"dinsos_kuburaya/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// =====================================
	// REGISTRASI - TANPA AUTH (PUBLIC)
	// =====================================
	api.POST("/users", controllers.CreateUser) // ← DIPINDAH KE SINI (tidak pakai middleware)

	// =====================================
	// RUTE YANG MEMERLUKAN AUTENTIKASI
	// =====================================
	users := api.Group("/users")
	users.Use(middleware.RequireAuth()) // ← Middleware hanya untuk rute di bawah ini
	{
		users.GET("", controllers.GetUsers)          // GET /api/users
		users.GET("/:id", controllers.GetUserByID)   // GET /api/users/:id
		users.PUT("/:id", controllers.UpdateUser)    // PUT /api/users/:id
		users.DELETE("/:id", controllers.DeleteUser) // DELETE /api/users/:id
	}
}
