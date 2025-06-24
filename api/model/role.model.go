package model

import "time"

// Role represents a user role
type Role struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Name      string       `gorm:"size:50;not null;unique" json:"name"`
	Users     []User       `gorm:"foreignKey:RoleID" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
