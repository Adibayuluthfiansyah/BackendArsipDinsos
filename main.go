package main

import (
	"dinsos_kuburaya/config"
	"dinsos_kuburaya/middleware"
	"dinsos_kuburaya/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Connect DB
	config.ConnectDatabase()

	// Setup Gin
	r := gin.Default()

	// ✅ TAMBAHKAN INI: CORS Middleware HARUS PERTAMA!
	r.Use(middleware.CORSMiddleware())

	// Register routes
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.DocumentRoutes(r)
	routes.NotificationRoutes(r)

	// Run server
	r.Run(":8080")
}
