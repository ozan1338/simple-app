package domain

import (
    "motivational-reminder/domain/repository"
    "motivational-reminder/domain/usecase"
    "motivational-reminder/infrastructure"
)

type Domain struct {
    QuoteUsecase usecase.QuoteUsecase
}

func ConstructDomain() Domain {
    databaseConn := infrastructure.NewDatabaseConnection()

    databaseRepository := repository.NewDatabaseRepository(databaseConn)

    quoteUsecase := usecase.NewQuoteUsecase(databaseRepository)

    return Domain{
        QuoteUsecase: &quoteUsecase,
    }
}
