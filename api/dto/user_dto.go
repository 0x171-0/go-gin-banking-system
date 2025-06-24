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
