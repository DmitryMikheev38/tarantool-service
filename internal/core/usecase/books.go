package usecase

import (
	"taran/internal/core/domain"
	"taran/internal/core/interfaces"
)

type BooksUseCase struct {
	repository        interfaces.BooksRepository
	authorsRepository interfaces.AuthorsRepository
}

func NewBooksUseCase(repository interfaces.BooksRepository, authorsRepository interfaces.AuthorsRepository) *BooksUseCase {
	return &BooksUseCase{
		repository:        repository,
		authorsRepository: authorsRepository,
	}
}

func (uc *BooksUseCase) Create(book *domain.Book) error {
	if err := uc.repository.Save(book); err != nil {
		return err
	}

	if book.Authors != nil {
		for _, author := range book.Authors {
			if err := uc.authorsRepository.Save(author); err != nil {
				return err
			}

			if err := uc.repository.SaveAuthorRelationship(book.ID, author.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *BooksUseCase) GetList() ([]*domain.Book, error) {
	return uc.repository.GetList()
}
