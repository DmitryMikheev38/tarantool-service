package domain

type Book struct {
	ID          string
	Title       string
	Description string
	Authors     []*Author
}
