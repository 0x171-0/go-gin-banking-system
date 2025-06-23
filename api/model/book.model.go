package model

import (
	"time"
)

// Book represents a book in the system
type Book struct {
	// The unique identifier of the book
	ID        uint      `gorm:"primaryKey" json:"id"`
	// The title of the book
	Title     string    `gorm:"not null" json:"title" example:"Go Programming"`
	// The author of the book
	Author    string    `gorm:"not null" json:"author" example:"John Doe"`
	// Creation time
	CreatedAt time.Time `json:"created_at"`
	// Update time
	UpdatedAt time.Time `json:"updated_at"`
}
