package model

import "time"

// Cart represents a user's shopping cart
type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"unique" json:"user_id"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	BookID    uint      `json:"book_id"`
	Book      Book      `json:"book"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
