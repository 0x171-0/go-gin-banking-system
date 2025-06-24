package model

import (
	"time"
)

// TransactionType represents the type of transaction
type TransactionType string

const (
	TransactionTypeTransfer TransactionType = "transfer"
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeWithdraw TransactionType = "withdraw"
)

// TransactionStatus represents the status of transaction
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusVerified  TransactionStatus = "verified"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusCanceled  TransactionStatus = "canceled"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID            uint             `gorm:"primaryKey" json:"id"`
	Amount        float64          `gorm:"type:decimal(20,8);not null" json:"amount"`
	Type          TransactionType  `gorm:"size:20;not null" json:"type"`
	Status        TransactionStatus `gorm:"size:20;not null;default:'pending'" json:"status"`
	Description   string           `gorm:"type:text" json:"description"`
	FromAccountID *uint            `json:"from_account_id,omitempty"`
	ToAccountID   *uint            `json:"to_account_id,omitempty"`
	FromAccount   *Account         `gorm:"foreignKey:FromAccountID" json:"from_account,omitempty"`
	ToAccount     *Account         `gorm:"foreignKey:ToAccountID" json:"to_account,omitempty"`
	Verifications []TransactionVerification `gorm:"foreignKey:TransactionID" json:"verifications,omitempty"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}
