package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-gin-template/api/config"
	"go-gin-template/api/model"
)

type BookRepository interface {
	FindAll() []model.Book
	FindByID(id int) (model.Book, error)
	Create(book model.Book) model.Book
}

type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) FindAll() []model.Book {
	var books []model.Book
	config.DB.Find(&books)
	return books
}

func (r *bookRepository) FindByID(id int) (model.Book, error) {
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

func (r *bookRepository) Create(book model.Book) model.Book {
	config.DB.Create(&book)

	// Cache the new book
	ctx := context.Background()
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, _ := json.Marshal(book)
	config.Redis.Set(ctx, key, bookJSON, 1*time.Hour)

	return book
}
