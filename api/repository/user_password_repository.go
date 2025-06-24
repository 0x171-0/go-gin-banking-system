package repository

import (
	"go-gin-template/api/model"

	"gorm.io/gorm"
)

type UserPasswordRepository interface {
	Create(password *model.UserPassword) error
	FindActiveByUserID(userID uint) (*model.UserPassword, error)
	DeactivateAll(userID uint) error
}

type userPasswordRepository struct {
	db *gorm.DB
}

func NewUserPasswordRepository(db *gorm.DB) UserPasswordRepository {
	return &userPasswordRepository{db: db}
}

func (r *userPasswordRepository) Create(password *model.UserPassword) error {
	return r.db.Create(password).Error
}

func (r *userPasswordRepository) FindActiveByUserID(userID uint) (*model.UserPassword, error) {
	var password model.UserPassword
	if err := r.db.Where("user_id = ? AND is_active = ?", userID, true).First(&password).Error; err != nil {
		return nil, err
	}
	return &password, nil
}

func (r *userPasswordRepository) DeactivateAll(userID uint) error {
	return r.db.Model(&model.UserPassword{}).Where("user_id = ?", userID).Update("is_active", false).Error
}
