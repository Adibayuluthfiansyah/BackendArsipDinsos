package controllers

import (
	"net/http"
	"strings"

	"dinsos_kuburaya/config"
	"dinsos_kuburaya/models"

	"github.com/gin-gonic/gin"
)

// Logout menghapus token dari tabel secret_tokens
func Logout(c *gin.Context) {
	// ✅ Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan"})
		return
	}

	// Extract token (format: "Bearer <token>")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Format token tidak valid"})
		return
	}

	db := config.DB

	// ✅ Cari dan hapus token berdasarkan token string
	var secretToken models.SecretToken
	if err := db.Where("token = ?", tokenString).First(&secretToken).Error; err != nil {
		// Token tidak ditemukan di database (mungkin sudah expired atau tidak valid)
		// Tapi tetap return success karena tujuan logout tercapai
		c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
		return
	}

	// Hapus token dari database
	if err := db.Delete(&secretToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout berhasil",
	})
}
