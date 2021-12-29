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
	author.ID = id
	return err
}

func (r *AuthorsRepository) GetList(limit, offset uint32, booksLimit int) ([]*domain.Author, error) {
	var result []*domain.Author
	err := r.db.SelectTyped("authors", "books_count", offset, limit, tarantool.IterGt, []interface{}{booksLimit}, &result)
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

func (r *AuthorsRepository) Update(author *domain.Author) error {
	_, err := r.db.Replace(
		"authors",
		[]interface{}{author.ID, author.Name, author.BooksCount},
	)
	return err
}
