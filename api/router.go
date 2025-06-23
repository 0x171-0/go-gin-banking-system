package api

import (
	"go-gin-template/api/handler"
	"go-gin-template/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	bookService := service.NewBookService()
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBook)
	r.POST("/books", bookHandler.CreateBook)

	return r
}
