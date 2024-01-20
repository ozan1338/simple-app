package domain

import (
    "system-gift-point/domain/repository"
    "system-gift-point/domain/usecase"
    "system-gift-point/infrastructure"
)

type Domain struct {
    ActionUsecase usecase.ActionUsecase
}

func ConstructDomain() Domain {
    databaseConn := infrastructure.NewDatabaseConnection()

    databaseRepository := repository.NewDatabaseRepository(databaseConn)

    actionUsecase := usecase.NewActionUsecase(databaseRepository)

    return Domain{
        ActionUsecase: &actionUsecase,
    }
}
