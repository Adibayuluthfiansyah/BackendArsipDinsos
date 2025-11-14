package controllers

import (
	"dinsos_kuburaya/config"
	"dinsos_kuburaya/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	// AMBIL DARI CONTEXT (BUKAN HARDCODE)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
		return
	}
	userID := userIDInterface.(string) // Konversi ke string

	var notifications []models.Notification
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil notifikasi"})
		return
	}

	var unreadCount int64
	config.DB.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&unreadCount)

	c.JSON(http.StatusOK, gin.H{
		"notifications": notifications,
		"unread_count":  unreadCount,
	})
}

func MarkNotificationAsRead(c *gin.Context) {
	// AMBIL DARI CONTEXT
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
		return
	}
	userID := userIDInterface.(string)
	notifID := c.Param("id")

	var notification models.Notification
	// Pastikan notifikasi ini milik user yang login
	if err := config.DB.Where("id = ? AND user_id = ?", notifID, userID).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notifikasi tidak ditemukan"})
		return
	}

	notification.IsRead = true
	config.DB.Save(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Notifikasi ditandai sebagai dibaca"})
}
