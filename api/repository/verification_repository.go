package repository

import (
	"go-gin-template/api/model"
	"gorm.io/gorm"
)

type VerificationRepository interface {
	Create(verification *model.TransactionVerification) error
	FindByID(verificationID uint) (*model.TransactionVerification, error)
	FindByTransactionID(transactionID uint) ([]*model.TransactionVerification, error)
	Update(verification *model.TransactionVerification) error
	UpdateStatus(verificationID uint, status model.VerificationStatus) error
	FindActiveByTransactionID(transactionID uint) (*model.TransactionVerification, error)
}

type verificationRepository struct {
	db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) VerificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) Create(verification *model.TransactionVerification) error {
	return r.db.Create(verification).Error
}

func (r *verificationRepository) FindByID(verificationID uint) (*model.TransactionVerification, error) {
	var verification model.TransactionVerification
	err := r.db.Where("id = ?", verificationID).First(&verification).Error
	return &verification, err
}

func (r *verificationRepository) FindByTransactionID(transactionID uint) ([]*model.TransactionVerification, error) {
	var verifications []*model.TransactionVerification
	err := r.db.Where("transaction_id = ?", transactionID).Find(&verifications).Error
	return verifications, err
}

func (r *verificationRepository) Update(verification *model.TransactionVerification) error {
	return r.db.Save(verification).Error
}

func (r *verificationRepository) UpdateStatus(verificationID uint, status model.VerificationStatus) error {
	return r.db.Model(&model.TransactionVerification{}).
		Where("id = ?", verificationID).
		Update("status", status).Error
}

func (r *verificationRepository) FindActiveByTransactionID(transactionID uint) (*model.TransactionVerification, error) {
	var verification model.TransactionVerification
	err := r.db.Where("transaction_id = ? AND status = ?", transactionID, model.VerificationStatusPending).
		Order("created_at DESC").
		First(&verification).Error
	return &verification, err
}