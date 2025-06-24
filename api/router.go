package api

import (
	"go-gin-template/api/handler"
	"go-gin-template/api/middleware"
	"go-gin-template/api/repository"
	"go-gin-template/api/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// Use logger and recovery middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Use custom error handler
	r.Use(middleware.ErrorHandler())

	// Book endpoints
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// Public book endpoints
	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBook)

	// Protected book endpoints
	bookGroup := r.Group("/books")
	bookGroup.Use(middleware.AuthMiddleware())
	{
		bookGroup.POST("", bookHandler.CreateBook)
	}

	// User endpoints
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Public user endpoints
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}

	// Protected user endpoints
	protectedUserGroup := r.Group("/users")
	protectedUserGroup.Use(middleware.AuthMiddleware())
	{
		// These endpoints require authentication and ownership/admin rights
		protectedUserGroup.GET("/:id", middleware.OwnerOrAdminAuthMiddleware(), userHandler.GetProfile)
		protectedUserGroup.PUT("/:id", middleware.OwnerOrAdminAuthMiddleware(), userHandler.UpdateProfile)
	}

	return r
}
