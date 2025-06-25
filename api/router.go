package api

import (
	"go-gin-template/api/config"
	"go-gin-template/api/handler"
	"go-gin-template/api/middleware"
	"go-gin-template/api/repository"
	"go-gin-template/api/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Initialize repositories
	bookRepo := repository.NewBookRepository(config.DB)
	userRepo := repository.NewUserRepository(config.DB)
	accountRepo := repository.NewAccountRepository(config.DB)
	passwordRepo := repository.NewUserPasswordRepository(config.DB)
	transactionRepo := repository.NewTransactionRepository(config.DB)
	r := gin.Default()

	// Use recovery middleware
	r.Use(gin.Recovery())

	// Use custom error handler
	r.Use(middleware.ErrorHandler())

	// Book endpoints
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBook)
	r.POST("/books", middleware.AuthMiddleware(), bookHandler.CreateBook)
	r.PUT("/books/:id", middleware.AuthMiddleware(), bookHandler.UpdateBook)
	r.DELETE("/books/:id", middleware.AuthMiddleware(), bookHandler.DeleteBook)

	// User endpoints
	accountService := service.NewAccountService(accountRepo, transactionRepo)

	userService := service.NewUserService(userRepo, passwordRepo, accountService)
	userHandler := handler.NewUserHandler(userService)
	users := r.Group("/users")
	{
		users.POST("/login", userHandler.Login)
		users.POST("/register", userHandler.Register)
		users.GET("/:id", middleware.AuthMiddleware(), userHandler.GetProfile)
		users.PUT("/:id", middleware.AuthMiddleware(), userHandler.UpdateProfile)
	}

	// Account endpoints
	accountHandler := handler.NewAccountHandler(accountService)
	accounts := r.Group("/accounts", middleware.AuthMiddleware())
	{
		accounts.POST("", accountHandler.CreateAccount)
		accounts.GET("", accountHandler.GetAccounts)
		accounts.POST("/:id/deposit", middleware.AccountOwnershipGuard(), accountHandler.Deposit)
		accounts.POST("/:id/withdraw", middleware.AccountOwnershipGuard(), accountHandler.Withdraw)
		accounts.POST("/:id/transfer", middleware.AccountOwnershipGuard(), accountHandler.Transfer)
	}

	return r
}
