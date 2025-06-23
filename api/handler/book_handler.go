package handler

import (
	"github.com/white/go-gin-template/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.bookService.GetAll())
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.bookService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到書籍"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := h.bookService.Create(req.Title, req.Author)
	c.JSON(http.StatusCreated, book)
}
