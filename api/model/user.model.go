package model

import "time"

// User represents a user in the system
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // Password is not included in JSON
	Name      string    `gorm:"not null" json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Role      string    `gorm:"default:'user'" json:"role"` // 'user' or 'admin'
	Orders    []Order   `json:"orders,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
