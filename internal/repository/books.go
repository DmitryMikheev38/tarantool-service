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
	_, err := r.db.Insert("books", []interface{}{id, book.Title, book.Description})
	book.ID = id
	return err
}

func (r *BooksRepository) SaveAuthorRelationship(bookID, authorID string) error {
	_, err := r.db.Insert("books_authors", []interface{}{authorID, bookID})
	return err
}

func (r *BooksRepository) GetList() ([]*domain.Book, error) {
	var result []*domain.Book

	err := r.db.SelectTyped("books", "pk", 0, 10, tarantool.IterAll, []interface{}{}, &result)
	return nil, err
}
