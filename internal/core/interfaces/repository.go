package interfaces

import "taran/internal/core/domain"

type BooksRepository interface {
	Save(book *domain.Book) error
	Delete(book *domain.Book) error
	Update(book *domain.Book) error
	GetList(limit, offset uint32) ([]*domain.Book, error)
	GetByAuthorList(authorID string, limit, offset uint32) ([]*domain.Book, error)
}

type AuthorsRepository interface {
	Save(author *domain.Author) error
	GetList(limit, offset uint32, booksLimit int) ([]*domain.Author, error)
	GetByID(id string) (*domain.Author, error)
	GetByBookID(id string) (*domain.Author, error)
	Update(author *domain.Author) error
}
