package model

import "time"

// User represents a user in the system
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	RoleID    *uint     `gorm:"column:role_id" json:"role_id,omitempty"`
	Role      *Role     `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Password  *UserPassword `gorm:"foreignKey:UserID" json:"password,omitempty"`
	Accounts  []Account `gorm:"foreignKey:UserID" json:"accounts,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
