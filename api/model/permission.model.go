package model

import "time"

// Permission represents a system permission
type Permission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null;unique" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Roles       []Role    `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
