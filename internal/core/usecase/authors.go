package usecase

import (
	"taran/internal/core/domain"
	"taran/internal/core/interfaces"
)

type AuthorsUseCase struct {
	repository      interfaces.AuthorsRepository
	booksRepository interfaces.BooksRepository
}

func NewAuthorsUseCase(repository interfaces.AuthorsRepository, booksRepository interfaces.BooksRepository) *AuthorsUseCase {
	return &AuthorsUseCase{
		repository:      repository,
		booksRepository: booksRepository,
	}
}

func (uc *AuthorsUseCase) Create(author *domain.Author) error {
	if err := uc.repository.Save(author); err != nil {
		return err
	}

	return nil
}

func (uc *AuthorsUseCase) GetList(booksLimit int, limit, offset uint32) ([]*domain.Author, error) {
	return uc.repository.GetList(limit, offset, booksLimit)
}

func (uc *AuthorsUseCase) AddBook(book *domain.Book) error {
	author, err := uc.repository.GetByID(book.AuthorID)

	if err != nil {
		return err
	}

	if author == nil {
		return nil
	}

	if err := uc.booksRepository.Save(book); err != nil {
		return err
	}

	author.BooksCount++

	return uc.repository.Update(author)
}

func (uc *AuthorsUseCase) GetBooksList(authorID string, limit, offset uint32) ([]*domain.Book, error) {
	return uc.booksRepository.GetByAuthorList(authorID, limit, offset)
}
