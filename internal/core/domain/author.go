package domain

type Author struct {
	ID         string
	Name       string
	BooksCount int
	Books      []*Book
}
