package model

// Book represents a book in the system
type Book struct {
	// The unique identifier of the book
	ID int `json:"id" example:"1"`
	// The title of the book
	Title string `json:"title" example:"Go Programming"`
	// The author of the book
	Author string `json:"author" example:"John Doe"`
}
