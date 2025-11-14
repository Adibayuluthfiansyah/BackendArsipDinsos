package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"dinsos_kuburaya/config"
	"dinsos_kuburaya/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RequireAuth adalah middleware untuk rute yang WAJIB login
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak ditemukan",
			})
			return
		}

		// Cek format "Bearer <token>"
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Format token salah. Gunakan: Bearer <token>",
			})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validasi signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak terduga: %v", token.Header["alg"])
			}

			secret := os.Getenv("SECRET_TOKEN")
			if secret == "" {
				return nil, fmt.Errorf("SECRET_TOKEN tidak dikonfigurasi")
			}

			return []byte(secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token tidak valid",
				"detail": err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak valid atau claims rusak",
			})
			return
		}

		// Cek kedaluwarsa
		exp, ok := claims["exp"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak memiliki expiration",
			})
			return
		}

		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token kedaluwarsa",
			})
			return
		}

		// Ambil user dari database
		sub, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Subject token tidak valid",
			})
			return
		}

		var user models.User
		if err := config.DB.First(&user, "id = ?", sub).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User tidak ditemukan",
			})
			return
		}
		c.Set("userID", user.ID)
		c.Set("userRole", user.Role)
		c.Next()
	}
}

// AdminOnly adalah middleware yang memeriksa apakah user adalah admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role.(string) != "admin" {
			log.Printf("❌ [ADMIN] Access denied for role: %v", role)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Akses ditolak. Memerlukan hak admin.",
			})
			return
		}
		c.Next()
	}
}

// OptionalAuth adalah middleware untuk rute registrasi
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString != "" && strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_TOKEN")), nil
			})

			if err == nil && token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)
				exp, expOk := claims["exp"].(float64)
				sub, subOk := claims["sub"].(string)

				if ok && expOk && subOk && float64(time.Now().Unix()) <= exp {
					var user models.User
					if err := config.DB.First(&user, "id = ?", sub).Error; err == nil {
						c.Set("userID", user.ID)
						c.Set("userRole", user.Role)
					}
				}
			}
		}
		c.Next()
	}
}
