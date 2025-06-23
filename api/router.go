package api

import (
	"go-gin-template/api/handler"
	"go-gin-template/api/repository"
	"go-gin-template/api/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	bookRepo := repository.NewBookRepository()
bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBook)
	r.POST("/books", bookHandler.CreateBook)

	return r
}
