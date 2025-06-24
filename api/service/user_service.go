package service

import (
	"errors"

	"go-gin-template/api/dto"
	"go-gin-template/api/model"
	"go-gin-template/api/repository"
	"go-gin-template/api/util"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetUserByID(id uint) (*dto.UserResponse, error)
	UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	userRepo         repository.UserRepository
	passwordRepo     repository.UserPasswordRepository
	accountService   AccountService
}

func NewUserService(userRepo repository.UserRepository, passwordRepo repository.UserPasswordRepository, accountService AccountService) UserService {
	return &userService{
		userRepo:       userRepo,
		passwordRepo:   passwordRepo,
		accountService: accountService,
	}
}

func (s *userService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	// Check if user already exists
	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &model.User{
		Email:   req.Email,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}

	// Create user first
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Create password
	password := &model.UserPassword{
		UserID:         user.ID,
		HashedPassword: string(hashedPassword),
		IsActive:      true,
	}

	// Set password on user for response
	user.Password = password

	// Save password
	if err := s.passwordRepo.DeactivateAll(user.ID); err != nil {
		return nil, err
	}

	if err := s.passwordRepo.Create(password); err != nil {
		return nil, err
	}

	// Create default account for new user
	_, err = s.accountService.CreateDefaultAccount(user.ID)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Get active password
	activePassword, err := s.passwordRepo.FindActiveByUserID(user.ID)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Set active password on user
	user.Password = activePassword

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.HashedPassword), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Get role name for token
	roleName := "user"
	if user.Role != nil {
		roleName = user.Role.Name
	}

	// Generate JWT token
	token, err := util.GenerateToken(user.ID, roleName)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		User:  *toUserResponse(user),
		Token: token,
	}, nil
}

func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func (s *userService) UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Address != "" {
		user.Address = req.Address
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

// Helper function to convert User model to UserResponse DTO
func toUserResponse(user *model.User) *dto.UserResponse {
	roleName := "user"
	if user.Role != nil {
		roleName = user.Role.Name
	}

	return &dto.UserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
		Role:    roleName,
	}
}
