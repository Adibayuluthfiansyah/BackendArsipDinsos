package main

import (
	"log"

	"dinsos_kuburaya/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"dinsos_kuburaya/config"
	"dinsos_kuburaya/models"
	"dinsos_kuburaya/routes"
)

func main() {
	// ✅ Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	r := gin.Default()

	// Connect Database
	config.ConnectDatabase()

	// Auto Migrate
	if err := config.DB.AutoMigrate(&models.User{}, &models.Document{}, &models.SecretToken{}); err != nil {
		log.Fatal("Gagal migrasi tabel:", err)
	}

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimiter())

	// ✅ 3. Baru daftarkan routes
	routes.LoginRoutes(r)
	routes.UserRoutes(r)
	routes.DocumentRoutes(r)
	routes.LogoutRoutes(r)

	// ✅ 4. Jalankan server
	r.Run(":8080")
}
