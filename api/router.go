package api

import (
	"go-gin-template/api/handler"
	"go-gin-template/api/middleware"
	"go-gin-template/api/repository"
	"go-gin-template/api/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(userRepo repository.UserRepository, bookRepo repository.BookRepository) *gin.Engine {
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
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.GET("/users/:id", middleware.AuthMiddleware(), userHandler.GetProfile)
	r.PUT("/users/:id", middleware.AuthMiddleware(), userHandler.UpdateProfile)

	return r
}
