package repository

import (
	"go-gin-template/api/model"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *model.Account) error
	FindByID(id uint) (*model.Account, error)
	FindByUserID(userID uint) ([]*model.Account, error)
	FindDefaultByUserID(userID uint) (*model.Account, error)
	Update(account *model.Account) error
	GetDB() *gorm.DB
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) FindByID(id uint) (*model.Account, error) {
	var account model.Account
	err := r.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindByUserID(userID uint) ([]*model.Account, error) {
	var accounts []*model.Account
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) FindDefaultByUserID(userID uint) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("user_id = ? AND is_default = ?", userID, true).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) Update(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) GetDB() *gorm.DB {
	return r.db
}
