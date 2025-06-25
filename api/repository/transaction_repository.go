package repository

import (
	"go-gin-template/api/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	FindByID(transactionID uint) (*model.Transaction, error)
	Update(transaction *model.Transaction) error
	UpdateStatus(transactionID uint, status model.TransactionStatus) error
	GetDB() *gorm.DB
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindByID(transactionID uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Preload("FromAccount").Preload("ToAccount").
		Where("id = ?", transactionID).First(&transaction).Error
	return &transaction, err
}

func (r *transactionRepository) Update(transaction *model.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) UpdateStatus(transactionID uint, status model.TransactionStatus) error {
	return r.db.Model(&model.Transaction{}).
		Where("id = ?", transactionID).
		Update("status", status).Error
}

func (r *transactionRepository) GetDB() *gorm.DB {
	return r.db
}