package model

import "time"

// UserPassword represents user's password and security information
type UserPassword struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;unique" json:"user_id"`
	HashedPassword string    `gorm:"size:255;not null" json:"-"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	HMACSecret     string    `gorm:"size:255;column:hmac_secret" json:"-"`
	User           User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
