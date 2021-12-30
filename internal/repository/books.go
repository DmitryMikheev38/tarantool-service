package repository

import (
	"github.com/google/uuid"
	"github.com/tarantool/go-tarantool"
	"taran/internal/core/domain"
)

type BooksRepository struct {
	db tarantool.Connector
}

func NewBooksRepository(connection tarantool.Connector) *BooksRepository {
	return &BooksRepository{
		db: connection,
	}
}

func (r *BooksRepository) Save(book *domain.Book) error {
	id := uuid.New().String()
	_, err := r.db.Insert("books", []interface{}{id, book.AuthorID, book.Title, book.Description})
	if err != nil {
		return err
	}
	book.ID = id

	return nil
}

func (r *BooksRepository) GetList(limit, offset uint32) ([]*domain.Book, error) {
	var result []*domain.Book
	err := r.db.SelectTyped("books", "pk", offset, limit, tarantool.IterAll, []interface{}{}, &result)

	return result, err
}

func (r *BooksRepository) GetByAuthorList(authorID string, limit, offset uint32) ([]*domain.Book, error) {
	var books []*domain.Book

	err := r.db.SelectTyped("books", "author_idx", offset, limit, tarantool.IterEq, []interface{}{authorID}, &books)
	if err != nil {
		return nil, err
	}

	return books, err
}

func (r *BooksRepository) Delete(book *domain.Book) error {
	_, err := r.db.Delete("books", "pk", []interface{}{book.ID})

	return err
}

func (r *BooksRepository) Update(book *domain.Book) error {
	_, err := r.db.Replace("books", []interface{}{book.ID, book.AuthorID, book.Title, book.Description})

	return err
}
