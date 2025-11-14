package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    string    `gorm:"type:char(36);not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Message   string    `gorm:"type:varchar(255)" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	Link      string    `gorm:"type:varchar(255);default:null" json:"link"`
	CreatedAt time.Time `json:"created_at"`
}

// Generate UUID sebelum disimpan
func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.NewString()
	return
}
