package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-gin-template/api/config"
	"go-gin-template/api/model"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) GetAll() []model.Book {
	var books []model.Book
	config.DB.Find(&books)
	return books
}

func (s *BookService) GetByID(id int) (model.Book, error) {
	ctx := context.Background()
	// Try to get from Redis first
	key := fmt.Sprintf("book:%d", id)
	data, err := config.Redis.Get(ctx, key).Result()
	if err == nil {
		// Found in cache
		var book model.Book
		if err := json.Unmarshal([]byte(data), &book); err == nil {
			return book, nil
		}
	}

	// Not found in cache or unmarshal error, query database
	var book model.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return model.Book{}, errors.New("not found")
	}

	// Store in Redis
	bookJSON, _ := json.Marshal(book)
	config.Redis.Set(ctx, key, bookJSON, 1*time.Hour)

	return book, nil
}

func (s *BookService) Create(title, author string) model.Book {
	book := model.Book{
		Title:  title,
		Author: author,
	}

	config.DB.Create(&book)

	// Cache the new book
	ctx := context.Background()
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, _ := json.Marshal(book)
	config.Redis.Set(ctx, key, bookJSON, 1*time.Hour)

	return book
}
