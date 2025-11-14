package routes

import (
	"dinsos_kuburaya/controllers"
	"dinsos_kuburaya/middleware" // <-- IMPORT

	"github.com/gin-gonic/gin"
)

func DocumentRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Buat grup baru untuk dokumen
	docs := api.Group("/documents")
	docs.Use(middleware.RequireAuth()) // <-- LINindungi semua rute dokumen
	{
		docs.POST("", controllers.CreateDocument)
		docs.GET("", controllers.GetDocuments)
		docs.GET("/:id", controllers.GetDocumentByID)
		docs.PUT("/:id", controllers.UpdateDocument)
		docs.DELETE("/:id", controllers.DeleteDocument)
	}
}
