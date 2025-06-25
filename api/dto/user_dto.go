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

// TransferRequest represents the request body for transfer
type TransferRequest struct {
	Amount           float64 `json:"amount" binding:"required,gt=0" example:"100.50"`
	TargetAccountID uint    `json:"target_account_id" binding:"required"`
}

// TransferInitRequest represents the request body for initiating transfer
// Used by: POST /accounts/{id}/transfer/init
type TransferInitRequest struct {
	Amount           float64 `json:"amount" binding:"required,gt=0" example:"100.50"`
	TargetAccountID uint    `json:"target_account_id" binding:"required"`
	Description      string  `json:"description" example:"Payment for services"`
}

// VerificationRequest represents the request body for verification generation
// Used by: POST /verifications
type VerificationRequest struct {
	TransactionID uint   `json:"transaction_id" binding:"required"`
	Type          string `json:"type" binding:"required,oneof=email sms" example:"email"`
}

// VerificationVerifyRequest represents the request body for verification code verification
// Used by: POST /verifications/{id}/verify
type VerificationVerifyRequest struct {
	Code string `json:"code" binding:"required,len=6" example:"123456"`
}

// TransferInitResponse represents the response body for transfer initiation
// Used by: POST /accounts/{id}/transfer/init
type TransferInitResponse struct {
	TransactionID uint   `json:"transaction_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

// VerificationResponse represents the response body for verification operations
// Used by: POST /verifications, POST /verifications/{id}/verify
type VerificationResponse struct {
	ID        uint   `json:"id"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	ExpiresAt string `json:"expires_at,omitempty"`
}
