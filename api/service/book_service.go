package service

import (
	"go-gin-template/api/model"
	"go-gin-template/api/repository"
)

type BookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (s *BookService) GetAll() []model.Book {
	return s.bookRepo.FindAll()
}

func (s *BookService) GetByID(id int) (model.Book, error) {
	return s.bookRepo.FindByID(id)
}

func (s *BookService) Create(title, author string) model.Book {
	book := model.Book{
		Title:  title,
		Author: author,
	}
	return s.bookRepo.Create(book)
}
