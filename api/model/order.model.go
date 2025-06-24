package model

import "time"

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid     OrderStatus = "paid"
	OrderStatusShipped  OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// Order represents an order in the system
type Order struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	UserID        uint        `json:"user_id"`
	User          User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems    []OrderItem `json:"order_items"`
	TotalAmount   float64     `gorm:"not null" json:"total_amount"`
	Status        OrderStatus `gorm:"not null;default:'pending'" json:"status"`
	PaymentMethod string      `json:"payment_method"`
	PaymentID     string      `json:"payment_id"`
	ShippingAddr  string      `gorm:"type:text" json:"shipping_address"`
	TrackingNo    string      `json:"tracking_no"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	BookID    uint      `json:"book_id"`
	Book      Book      `json:"book"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"` // Price at the time of purchase
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
