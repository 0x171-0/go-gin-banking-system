package middleware

import (
	"go-gin-template/api/config"
	"go-gin-template/api/model"
	"go-gin-template/api/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the JWT token and sets user information in the context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the Authorization header has the correct format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Parse and validate the token
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in the context
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

// AdminAuthMiddleware verifies if the user has admin role
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// OwnerOrAdminAuthMiddleware verifies if the user is the owner of the resource or an admin
func OwnerOrAdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		role, _ := c.Get("userRole")
		resourceUserID := c.Param("id")

		// Allow access if user is admin or the owner of the resource
		if role != "admin" && resourceUserID != "" && userID != resourceUserID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AccountOwnershipGuard verifies if the user owns the account specified in the path parameter
func AccountOwnershipGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if user has admin role
		role, _ := c.Get("userRole")
		if role == "admin" {
			c.Next()
			return
		}

		// Get account ID from path parameter
		accountIDStr := c.Param("id")
		if accountIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
			c.Abort()
			return
		}

		accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			c.Abort()
			return
		}

		// Query the account to verify ownership
		var account model.Account
		if err := config.DB.Where("id = ? AND user_id = ?", accountID, userID).First(&account).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Account not found or access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
