package repository

import (
	"github.com/google/uuid"
	"github.com/tarantool/go-tarantool"
	"taran/internal/core/domain"
)

type AuthorsRepository struct {
	db tarantool.Connector
}

func NewAuthorsRepository(connection tarantool.Connector) *AuthorsRepository {
	return &AuthorsRepository{
		db: connection,
	}
}

func (r *AuthorsRepository) Save(author *domain.Author) error {
	id := uuid.New().String()
	_, err := r.db.Insert("authors", []interface{}{id, author.Name, author.BooksCount})
	if err != nil {
		return err
	}

	author.ID = id

	return nil
}

func (r *AuthorsRepository) GetList(limit, offset uint32, booksLimit int) ([]*domain.Author, error) {
	var result []*domain.Author
	err := r.db.SelectTyped("authors", "books_count_idx", offset, limit, tarantool.IterGe, []interface{}{booksLimit}, &result)
	return result, err
}

func (r *AuthorsRepository) GetByID(id string) (*domain.Author, error) {
	var result []*domain.Author
	err := r.db.SelectTyped("authors", "pk", 0, 1, tarantool.IterEq, []interface{}{id}, &result)
	if len(result) > 0 {
		return result[0], err
	}

	return nil, err
}

func (r *AuthorsRepository) GetByBookID(id string) (*domain.Author, error) {
	var books []*domain.Book
	err := r.db.SelectTyped("books", "pk", 0, 1, tarantool.IterEq, []interface{}{id}, &books)
	if len(books) == 0 || err != nil {
		return nil, err
	}
	var result []*domain.Author
	err = r.db.SelectTyped("authors", "pk", 0, 1, tarantool.IterEq, []interface{}{books[0].AuthorID}, &result)
	if len(books) == 0 || err != nil {
		return nil, err
	}

	return result[0], err
}

func (r *AuthorsRepository) Update(author *domain.Author) error {
	_, err := r.db.Replace(
		"authors",
		[]interface{}{author.ID, author.Name, author.BooksCount},
	)
	return err
}
