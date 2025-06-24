package handler

import (
	"go-gin-template/api/dto"
	"go-gin-template/api/middleware"
	"go-gin-template/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration information"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} object "Invalid input"
// @Router /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		c.Error(middleware.BadRequestError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary User login
// @Description Authenticate a user and return user information
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} object "Authentication failed"
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	response, err := h.userService.Login(&req)
	if err != nil {
		c.Error(middleware.UnauthorizedError())
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} object "Invalid user ID"
// @Failure 401 {object} object "Unauthorized"
// @Failure 403 {object} object "Forbidden"
// @Failure 404 {object} object "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid user ID"))
		return
	}

	// User ID has already been validated by OwnerOrAdminAuthMiddleware
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.Error(middleware.NotFoundError("User"))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UpdateUserRequest true "User info to update"
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} object "Invalid input"
// @Failure 401 {object} object "Unauthorized"
// @Failure 403 {object} object "Forbidden"
// @Failure 404 {object} object "User not found"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid user ID"))
		return
	}

	// User ID has already been validated by OwnerOrAdminAuthMiddleware
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		c.Error(middleware.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, user)
}
