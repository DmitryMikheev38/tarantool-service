package interfaces

import (
	"taran/internal/core/domain"
)

type BooksUseCase interface {
	GetList(limit, offset uint32) ([]*domain.Book, error)
	BulkDelete(books []*domain.Book) error
	Update(book *domain.Book) error
}

type AuthorsUseCase interface {
	Create(book *domain.Author) error
	GetList(booksLimit int, limit, offset uint32) ([]*domain.Author, error)
	AddBook(book *domain.Book) error
	GetBooksList(authorID string, limit, offset uint32) ([]*domain.Book, error)
}
