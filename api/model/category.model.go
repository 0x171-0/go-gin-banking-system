package model

import "time"

// Category represents a book category
type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;uniqueIndex" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Books       []Book    `gorm:"foreignKey:CategoryID" json:"books,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
