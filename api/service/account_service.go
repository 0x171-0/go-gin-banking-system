package service

import (
	"errors"
	"go-gin-template/api/dto"
	"go-gin-template/api/model"
	"go-gin-template/api/repository"
)

type AccountService interface {
	CreateAccount(userID uint, req *dto.CreateAccountRequest) (*dto.AccountResponse, error)
	GetUserAccounts(userID uint) ([]*dto.AccountResponse, error)
	Deposit(userID, accountID uint, amount float64) (*dto.AccountResponse, error)
	Withdraw(userID, accountID uint, amount float64) (*dto.AccountResponse, error)
	Transfer(userID uint, sourceAccountID uint, targetAccountID uint, amount float64) (*dto.AccountResponse, error)
	CreateDefaultAccount(userID uint) (*dto.AccountResponse, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) CreateAccount(userID uint, req *dto.CreateAccountRequest) (*dto.AccountResponse, error) {
	account := &model.Account{
		UserID: userID,
		Name:   req.Name,
	}

	if err := s.accountRepo.Create(account); err != nil {
		return nil, err
	}

	return toAccountResponse(account), nil
}

func (s *accountService) GetUserAccounts(userID uint) ([]*dto.AccountResponse, error) {
	accounts, err := s.accountRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.AccountResponse
	for _, account := range accounts {
		responses = append(responses, toAccountResponse(account))
	}
	return responses, nil
}

func (s *accountService) Deposit(userID, accountID uint, amount float64) (*dto.AccountResponse, error) {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.UserID != userID {
		return nil, errors.New("unauthorized access to account")
	}

	account.Balance += amount
	account.Nonce++

	if err := s.accountRepo.Update(account); err != nil {
		return nil, err
	}

	return toAccountResponse(account), nil
}

func (s *accountService) Withdraw(userID, accountID uint, amount float64) (*dto.AccountResponse, error) {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.UserID != userID {
		return nil, errors.New("unauthorized access to account")
	}

	if account.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	account.Balance -= amount
	account.Nonce++

	if err := s.accountRepo.Update(account); err != nil {
		return nil, err
	}

	return toAccountResponse(account), nil
}

func (s *accountService) Transfer(userID uint, sourceAccountID uint, targetAccountID uint, amount float64) (*dto.AccountResponse, error) {
	// Get source account
	sourceAccount, err := s.accountRepo.FindByID(sourceAccountID)
	if err != nil {
		return nil, err
	}

	// Verify ownership of source account
	if sourceAccount.UserID != userID {
		return nil, errors.New("unauthorized access to source account")
	}

	// Get target account
	targetAccount, err := s.accountRepo.FindByID(targetAccountID)
	if err != nil {
		return nil, err
	}

	// Check sufficient balance
	if sourceAccount.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	// Update source account balance
	sourceAccount.Balance -= amount
	sourceAccount.Nonce++

	// Update target account balance
	targetAccount.Balance += amount
	targetAccount.Nonce++

	// Use a transaction to ensure both updates succeed or both fail
	tx := s.accountRepo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(sourceAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(targetAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return toAccountResponse(sourceAccount), nil
}

func (s *accountService) CreateDefaultAccount(userID uint) (*dto.AccountResponse, error) {
	account := &model.Account{
		UserID:    userID,
		Name:      "Default Account",
		IsDefault: true,
	}

	if err := s.accountRepo.Create(account); err != nil {
		return nil, err
	}

	return toAccountResponse(account), nil
}

func toAccountResponse(account *model.Account) *dto.AccountResponse {
	return &dto.AccountResponse{
		ID:        account.ID,
		UserID:    account.UserID,
		Name:      account.Name,
		Balance:   account.Balance,
		IsDefault: account.IsDefault,
	}
}
