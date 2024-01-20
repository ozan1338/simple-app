package usecase

import (
    "favourite-book/domain/entity"
    "favourite-book/domain/repository"
)

type BookUsecase interface {
    GetBooks() ([]entity.Book, error)
    CreateBook(book entity.Book) error
}

type bookUsecaseImpl struct {
    databaseRepository repository.DatabaseRepository
}

func NewBookUsecase(databaseRepository repository.DatabaseRepository) bookUsecaseImpl {
    return bookUsecaseImpl{
        databaseRepository: databaseRepository,
    }
}

func (this *bookUsecaseImpl) GetBooks() ([]entity.Book, error) {
    var books []entity.Book
    if err := this.databaseRepository.Find(&books); err != nil {
        return nil, err
    }

    return books, nil
}

func (this *bookUsecaseImpl) CreateBook(book entity.Book) error {
    if err := this.databaseRepository.Create(&book); err != nil {
        return err
    }

    return nil
}
