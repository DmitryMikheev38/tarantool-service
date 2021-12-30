package domain

type Author struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	BooksCount int    `json:"booksCount"`
}
