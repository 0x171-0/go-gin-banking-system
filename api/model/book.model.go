package model

import (
	"time"
)

// Book represents a book in the system
type Book struct {
	// The unique identifier of the book
	ID          uint      `gorm:"primaryKey" json:"id"`
	// The title of the book
	Title       string    `gorm:"not null" json:"title" example:"Go Programming"`
	// The author of the book
	Author      string    `gorm:"not null" json:"author" example:"John Doe"`
	// The ISBN of the book
	ISBN        string    `gorm:"unique" json:"isbn"`
	// The price of the book
	Price       float64   `gorm:"not null" json:"price"`
	// The stock quantity
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	// Book description
	Description string    `gorm:"type:text" json:"description"`
	// Book cover image URL
	CoverURL    string    `json:"cover_url"`
	// Book publisher
	Publisher   string    `json:"publisher"`
	// Publication date
	PubDate     time.Time `json:"pub_date"`
	// Category ID
	CategoryID  uint      `json:"category_id"`
	// Category relationship
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	// Creation time
	CreatedAt   time.Time `json:"created_at"`
	// Update time
	UpdatedAt   time.Time `json:"updated_at"`
}
