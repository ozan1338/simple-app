package domain

import (
	"favourite-book/domain/repository"
	"favourite-book/domain/usecase"
	"favourite-book/infrastructure"
)

type Domain struct {
    BookUsecase usecase.BookUsecase
}

func ConstructDomain() Domain {
    databaseConn := infrastructure.NewDatabaseConnection()

    databaseRepository := repository.NewDatabaseRepository(databaseConn)

    bookUsecase := usecase.NewBookUsecase(databaseRepository)

    return Domain{
        BookUsecase: &bookUsecase,
    }
}