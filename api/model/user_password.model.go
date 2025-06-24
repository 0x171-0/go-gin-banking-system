package model

import "time"

// UserPassword represents user's password and security information
type UserPassword struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;index:idx_user_active,priority:1" json:"user_id"`
	HashedPassword string    `gorm:"size:255;not null" json:"-"`
	IsActive       bool      `gorm:"default:true;index:idx_user_active,priority:2" json:"is_active"`
	HMACSecret     string    `gorm:"size:255;column:hmac_secret" json:"-"`
	User           User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
