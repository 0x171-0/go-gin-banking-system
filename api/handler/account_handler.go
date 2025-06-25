package handler

import (
	"go-gin-template/api/dto"
	"go-gin-template/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserIDFromContext(c *gin.Context) uint {
	userID, _ := c.Get("userID")
	return userID.(uint)
}

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account for the authenticated user
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body dto.CreateAccountRequest true "Account creation request"
// @Success 201 {object} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	userID := getUserIDFromContext(c)
	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.accountService.CreateAccount(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// GetAccounts godoc
// @Summary Get user accounts
// @Description Get all accounts of the authenticated user
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} dto.AccountResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts [get]
func (h *AccountHandler) GetAccounts(c *gin.Context) {
	userID := getUserIDFromContext(c)

	accounts, err := h.accountService.GetUserAccounts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// Deposit godoc
// @Summary Deposit money
// @Description Deposit money to a specific account
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Account ID"
// @Param request body dto.TransactionRequest true "Deposit request"
// @Success 200 {object} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts/{id}/deposit [post]
func (h *AccountHandler) Deposit(c *gin.Context) {
	userID := getUserIDFromContext(c)
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	var req dto.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.accountService.Deposit(userID, uint(accountID), req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// Withdraw godoc
// @Summary Withdraw money
// @Description Withdraw money from a specific account
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Account ID"
// @Param request body dto.TransactionRequest true "Withdrawal request"
// @Success 200 {object} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts/{id}/withdraw [post]
func (h *AccountHandler) Withdraw(c *gin.Context) {
	userID := getUserIDFromContext(c)
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	var req dto.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.accountService.Withdraw(userID, uint(accountID), req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary Transfer money
// @Description Transfer money from one account to another
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Source Account ID"
// @Param request body dto.TransferRequest true "Transfer request"
// @Success 200 {object} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts/{id}/transfer [post]
func (h *AccountHandler) Transfer(c *gin.Context) {
	userID := getUserIDFromContext(c)
	sourceAccountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid source account ID"})
		return
	}

	var req dto.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.accountService.Transfer(userID, uint(sourceAccountID), req.TargetAccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary Initiate transfer with verification
// @Description Initiate a transfer that requires verification
// @Tags accounts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Source Account ID"
// @Param request body dto.TransferInitRequest true "Transfer initiation request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /accounts/{id}/transfer/init [post]
func (h *AccountHandler) InitiateTransfer(c *gin.Context) {
	userID := getUserIDFromContext(c)
	sourceAccountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid source account ID"})
		return
	}

	var req dto.TransferInitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.accountService.InitiateTransfer(userID, uint(sourceAccountID), req.TargetAccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transaction.ID,
		"status":         transaction.Status,
		"amount":         transaction.Amount,
		"message":        "Transfer initiated. Please generate verification code to complete the transfer.",
	})
}
