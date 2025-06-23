package service

import (
	"errors"
	"go-gin-template/model"
)

type BookService struct {
	books []model.Book
}

func NewBookService() *BookService {
	return &BookService{
		books: []model.Book{
			{ID: 1, Title: "Go 入門", Author: "小王"},
			{ID: 2, Title: "Gin 教學", Author: "小李"},
		},
	}
}

func (s *BookService) GetAll() []model.Book {
	return s.books
}

func (s *BookService) GetByID(id int) (model.Book, error) {
	for _, b := range s.books {
		if b.ID == id {
			return b, nil
		}
	}
	return model.Book{}, errors.New("not found")
}

func (s *BookService) Create(title, author string) model.Book {
	newBook := model.Book{
		ID:     len(s.books) + 1,
		Title:  title,
		Author: author,
	}
	s.books = append(s.books, newBook)
	return newBook
}
