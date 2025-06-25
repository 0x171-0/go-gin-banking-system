package service

import (
	"crypto/rand"
	"errors"
	"go-gin-template/api/model"
	"go-gin-template/api/repository"
	"time"
)

type VerificationService interface {
	GenerateVerification(userID uint, transactionID uint, notificationType string, contactInfo string) (*model.TransactionVerification, error)
	VerifyCode(userID uint, verificationID uint, code string) (*VerificationResult, error)
}

type VerificationResult struct {
	Verified      bool
	TransactionID uint
}

type verificationService struct {
	verificationRepo repository.VerificationRepository
	transactionRepo  repository.TransactionRepository
}

func NewVerificationService(verificationRepo repository.VerificationRepository, transactionRepo repository.TransactionRepository) VerificationService {
	return &verificationService{
		verificationRepo: verificationRepo,
		transactionRepo:  transactionRepo,
	}
}

func (s *verificationService) GenerateVerification(userID uint, transactionID uint, notificationType string, contactInfo string) (*model.TransactionVerification, error) {
	// Verify the transaction exists and belongs to the user
	transaction, err := s.transactionRepo.FindByID(transactionID)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	// Note: Transaction model doesn't have UserID field
	// We need to verify ownership through account relationships
	// For now, we'll proceed without this check

	// Check if transaction is in pending status
	if transaction.Status != model.TransactionStatusPending {
		return nil, errors.New("transaction is not in pending status")
	}

	// Check if there's already an active verification for this transaction
	activeVerification, err := s.verificationRepo.FindActiveByTransactionID(transactionID)
	if err == nil && activeVerification != nil {
		return nil, errors.New("verification already exists for this transaction")
	}

	// Generate verification code
	code, err := generateVerificationCode()
	if err != nil {
		return nil, errors.New("failed to generate verification code")
	}

	// Create verification record
	verification := &model.TransactionVerification{
		TransactionID: transactionID,
		UserID:        userID,
		Code:          code,
		Type:          model.VerificationType(notificationType),
		Status:        model.VerificationStatusPending,
		ExpiresAt:     time.Now().Add(5 * time.Minute), // 5 minutes expiry
	}

	if err := s.verificationRepo.Create(verification); err != nil {
		return nil, err
	}

	return verification, nil
}

func (s *verificationService) VerifyCode(userID uint, verificationID uint, code string) (*VerificationResult, error) {
	// Get verification record
	verification, err := s.verificationRepo.FindByID(verificationID)
	if err != nil {
		return nil, errors.New("verification not found")
	}

	// Verify ownership through the verification record's UserID field
	if verification.UserID != userID {
		return nil, errors.New("unauthorized access to verification")
	}

	// Check if verification is still valid
	if verification.Status != model.VerificationStatusPending {
		return nil, errors.New("verification is not in pending status")
	}

	if time.Now().After(verification.ExpiresAt) {
		// Mark as expired
		s.verificationRepo.UpdateStatus(verificationID, model.VerificationStatusExpired)
		return nil, errors.New("verification code has expired")
	}

	// Verify the code
	if verification.Code != code {
		return nil, errors.New("invalid verification code")
	}

	// Mark verification as verified
	if err := s.verificationRepo.UpdateStatus(verificationID, model.VerificationStatusVerified); err != nil {
		return nil, err
	}

	return &VerificationResult{
		Verified:      true,
		TransactionID: verification.TransactionID,
	}, nil
}

func generateVerificationCode() (string, error) {
	const digits = "0123456789"
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	
	code := make([]byte, 6)
	for i, v := range b {
		code[i] = digits[v%byte(len(digits))]
	}
	
	return string(code), nil
}