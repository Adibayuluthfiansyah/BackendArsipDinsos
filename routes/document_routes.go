package routes

import (
	"dinsos_kuburaya/controllers"

	"github.com/gin-gonic/gin"
)

func DocumentRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// ✅ Tanpa trailing slash
		api.POST("/documents", controllers.CreateDocument)       // → /api/documents
		api.GET("/documents", controllers.GetDocuments)          // → /api/documents
		api.GET("/documents/:id", controllers.GetDocumentByID)   // → /api/documents/:id
		api.PUT("/documents/:id", controllers.UpdateDocument)    // → /api/documents/:id
		api.DELETE("/documents/:id", controllers.DeleteDocument) // → /api/documents/:id
	}
}
