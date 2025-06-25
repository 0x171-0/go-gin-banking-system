package handler

import (
	"go-gin-template/api/dto"
	"go-gin-template/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VerificationHandler struct {
	verificationService service.VerificationService
	notificationService service.NotificationService
}

func NewVerificationHandler(verificationService service.VerificationService, notificationService service.NotificationService) *VerificationHandler {
	return &VerificationHandler{
		verificationService: verificationService,
		notificationService: notificationService,
	}
}

// @Summary Generate verification code for transaction
// @Description Generate and send verification code for a pending transaction
// @Tags verification
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body dto.VerificationRequest true "Verification request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /verifications [post]
func (h *VerificationHandler) GenerateVerification(c *gin.Context) {
	userID := getUserIDFromContext(c)

	var req dto.VerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verification, err := h.verificationService.GenerateVerification(userID, req.TransactionID, req.Type, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send notification
	if err := h.notificationService.SendVerificationCode(req.Type, "", verification.Code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"verification_id": verification.ID,
		"expires_at":      verification.ExpiresAt,
		"message":         "Verification code sent successfully",
	})
}

// @Summary Verify transaction code
// @Description Verify the verification code for a transaction
// @Tags verification
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Verification ID"
// @Param request body dto.VerificationVerifyRequest true "Verification verify request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /verifications/{id}/verify [post]
func (h *VerificationHandler) VerifyCode(c *gin.Context) {
	userID := getUserIDFromContext(c)
	verificationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification ID"})
		return
	}

	var req dto.VerificationVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.verificationService.VerifyCode(userID, uint(verificationID), req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"verified":       result.Verified,
		"transaction_id": result.TransactionID,
		"message":        "Verification successful",
	})
}