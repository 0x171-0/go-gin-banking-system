package service

import (
	"go-gin-template/api/model"
	"go-gin-template/api/repository"
)

type BookService interface {
	GetBooks() ([]model.Book, error)
	GetBook(id uint) (*model.Book, error)
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) GetBooks() ([]model.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *bookService) GetBook(id uint) (*model.Book, error) {
	return s.bookRepo.FindByID(id)
}

func (s *bookService) CreateBook(book *model.Book) error {
	return s.bookRepo.Create(book)
}

func (s *bookService) UpdateBook(book *model.Book) error {
	return s.bookRepo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.bookRepo.Delete(id)
}
