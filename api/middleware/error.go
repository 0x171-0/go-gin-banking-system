package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// AppError represents a custom error type for the application
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// ErrorHandler is a middleware that handles errors returned from routes
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute request handlers
		c.Next()

		// Only handle the error if one has been set
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *AppError
		var validationErrs validator.ValidationErrors

		// Handle different types of errors
		switch {
		case errors.As(err, &appErr):
			// Handle application specific errors
			c.JSON(appErr.Code, gin.H{
				"error": appErr.Message,
			})

		case errors.Is(err, gorm.ErrRecordNotFound):
			// Handle not found errors
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Resource not found",
			})

		case errors.As(err, &validationErrs):
			// Handle validation errors
			var errMsgs []string
			for _, e := range validationErrs {
				errMsgs = append(errMsgs, formatValidationError(e))
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errMsgs,
			})

		default:
			// Handle unknown errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}

// formatValidationError formats validator.ValidationErrors into readable messages
func formatValidationError(e validator.FieldError) string {
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must be at least " + e.Param() + " characters long"
	case "max":
		return field + " must not be longer than " + e.Param() + " characters"
	default:
		return field + " is invalid"
	}
}

// ValidationError represents validation errors
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NotFoundError creates a not found AppError
func NotFoundError(resource string) *AppError {
	return NewAppError(http.StatusNotFound, resource+" not found")
}

// BadRequestError creates a bad request AppError
func BadRequestError(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message)
}

// UnauthorizedError creates an unauthorized AppError
func UnauthorizedError() *AppError {
	return NewAppError(http.StatusUnauthorized, "Unauthorized access")
}

// ForbiddenError creates a forbidden AppError
func ForbiddenError() *AppError {
	return NewAppError(http.StatusForbidden, "Access forbidden")
}

// InternalServerError creates an internal server error AppError
func InternalServerError() *AppError {
	return NewAppError(http.StatusInternalServerError, "Internal server error")
}
