package dto

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Phone    string `json:"phone" example:"1234567890"`
	Address  string `json:"address" example:"123 Main St"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// UpdateUserRequest represents the request body for updating user profile
type UpdateUserRequest struct {
	Name    string `json:"name" example:"John Doe"`
	Phone   string `json:"phone" example:"1234567890"`
	Address string `json:"address" example:"123 Main St"`
}

// UserResponse represents the response body for user information
type UserResponse struct {
	ID      uint   `json:"id" example:"1"`
	Email   string `json:"email" example:"user@example.com"`
	Name    string `json:"name" example:"John Doe"`
	Phone   string `json:"phone" example:"1234567890"`
	Address string `json:"address" example:"123 Main St"`
	Role    string `json:"role" example:"user"`
}

// LoginResponse represents the response body for successful login
type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// CreateAccountRequest represents the request body for creating a new account
type CreateAccountRequest struct {
	Name string `json:"name" binding:"required" example:"Savings Account"`
}

// AccountResponse represents the response body for account information
type AccountResponse struct {
	ID        uint    `json:"id" example:"1"`
	UserID    uint    `json:"user_id" example:"1"`
	Name      string  `json:"name" example:"Savings Account"`
	Balance   float64 `json:"balance" example:"1000.50"`
	IsDefault bool    `json:"is_default" example:"true"`
}

// TransactionRequest represents the request body for deposit/withdrawal
type TransactionRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0" example:"100.50"`
}
