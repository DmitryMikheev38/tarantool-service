package domain

type Book struct {
	ID          string `json:"id"`
	AuthorID    string `json:"authorId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
