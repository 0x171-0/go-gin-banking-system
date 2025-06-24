package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go-gin-template/api/config"
	"go-gin-template/api/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]model.Book, error)
	FindByID(id uint) (*model.Book, error)
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) FindByID(id uint) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	// Try to cache the book
	ctx := context.Background()
	key := fmt.Sprintf("book:%d", id)
	bookJSON, _ := json.Marshal(book)
	config.Redis.Set(ctx, key, bookJSON, 1*time.Hour)
	return &book, nil
}

func (r *bookRepository) Create(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}
	// Cache the new book
	ctx := context.Background()
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, _ := json.Marshal(book)
	config.Redis.Set(ctx, key, bookJSON, 1*time.Hour)
	return nil
}

func (r *bookRepository) Update(book *model.Book) error {
	if err := r.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
