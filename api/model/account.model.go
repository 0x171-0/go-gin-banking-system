package model

import (
	"time"
)

// Account represents a user's account
type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Balance   float64   `gorm:"type:decimal(20,8);not null;default:0" json:"balance"`
	Nonce     int       `gorm:"not null;default:0" json:"nonce"`
	IsDefault bool      `gorm:"default:false" json:"is_default"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
