package interfaces

import (
	"taran/internal/core/domain"
	"taran/internal/core/domain/dto"
)

type BooksUseCase interface {
	Create(book *domain.Book) error
	GetList() ([]*domain.Book, error)
}

type AuthorsUseCase interface {
	Create(book *domain.Author) error
	GetList(params dto.GetListAuthorsParams) ([]*domain.Author, error)
	AddBook(authorID string, book *domain.Book) error
	GetBook() (*domain.Book, error)
}
