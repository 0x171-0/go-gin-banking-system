package dto

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	// The title of the book
	Title string `json:"title" example:"The Go Programming Language"`
	// The author of the book
	Author string `json:"author" example:"Alan A. A. Donovan"`
}
