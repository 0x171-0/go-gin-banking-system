package handler

import (
	"go-gin-template/api/dto"
	"go-gin-template/api/service"
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

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} model.Book
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.bookService.GetAll())
}

// GetBook godoc
// @Summary Get a book by ID
// @Description Get details of a specific book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Book
// @Failure 404 {object} object "找不到書籍"
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.bookService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到書籍"})
		return
	}
	c.JSON(http.StatusOK, book)
}


// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided information
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body dto.CreateBookRequest true "Book information"
// @Success 201 {object} model.Book
// @Failure 400 {object} object "Invalid input"
// @Failure 401 {object} object "Unauthorized"
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := h.bookService.Create(req.Title, req.Author)
	c.JSON(http.StatusCreated, book)
}
