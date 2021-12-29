package usecase

import (
	"taran/internal/core/domain"
	"taran/internal/core/domain/dto"
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

func (uc *AuthorsUseCase) GetList(dto dto.GetListAuthorsParams) ([]*domain.Author, error) {
	return uc.repository.GetList(dto.Limit, dto.Offset, dto.BooksLimit)
}

func (uc *AuthorsUseCase) AddBook(authorID string, book *domain.Book) error {

	author, err := uc.repository.GetByID(authorID)
	if err != nil {
		return err
	}
	if author == nil {
		return nil
	}

	if err := uc.booksRepository.Save(book); err != nil {
		return err
	}

	if err := uc.booksRepository.SaveAuthorRelationship(book.ID, authorID); err != nil {
		return err
	}

	author.BooksCount++
	return uc.repository.Update(author)
}
