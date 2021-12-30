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

func (uc *BooksUseCase) GetList(limit, offset uint32) ([]*domain.Book, error) {
	return uc.repository.GetList(limit, offset)
}

func (uc *BooksUseCase) BulkDelete(books []*domain.Book) error {
	for _, book := range books {
		author, err := uc.authorsRepository.GetByBookID(book.ID)
		if author == nil || err != nil {
			return err
		}
		if err := uc.repository.Delete(book); err != nil {
			return err
		}

		author.BooksCount--
		if err := uc.authorsRepository.Update(author); err != nil {
			return err
		}
	}

	return nil
}

func (uc *BooksUseCase) Update(book *domain.Book) error {
	return uc.repository.Update(book)
}
