package model

import (
	"time"
)

// VerificationType represents the type of verification
type VerificationType string

const (
	VerificationTypeEmail VerificationType = "email"
	VerificationTypeSMS   VerificationType = "sms"
	VerificationTypeOTP   VerificationType = "otp"
)

// VerificationStatus represents the status of verification
type VerificationStatus string

const (
	VerificationStatusPending   VerificationStatus = "pending"
	VerificationStatusVerified  VerificationStatus = "verified"
	VerificationStatusExpired   VerificationStatus = "expired"
	VerificationStatusCanceled  VerificationStatus = "canceled"
)

// TransactionVerification represents a verification record for a transaction
type TransactionVerification struct {
	ID            uint               `gorm:"primaryKey" json:"id"`
	TransactionID uint               `gorm:"not null" json:"transaction_id"`
	UserID        uint               `gorm:"not null" json:"user_id"`
	Code          string             `gorm:"size:6;not null" json:"-"`
	Type          VerificationType   `gorm:"size:20;not null" json:"type"`
	Status        VerificationStatus `gorm:"size:20;not null;default:'pending'" json:"status"`
	ExpiresAt     time.Time         `gorm:"not null" json:"expires_at"`
	AttemptCount  int               `gorm:"not null;default:0" json:"attempt_count"`
	MaxAttempts   int               `gorm:"not null;default:3" json:"max_attempts"`
	VerifiedAt    *time.Time        `json:"verified_at"`
	Transaction   Transaction       `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	User          User             `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}
